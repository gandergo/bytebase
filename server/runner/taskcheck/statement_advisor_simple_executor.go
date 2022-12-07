package taskcheck

import (
	"context"
	"encoding/json"

	"github.com/bytebase/bytebase/api"
	"github.com/bytebase/bytebase/common"
	"github.com/bytebase/bytebase/plugin/advisor"
	advisorDB "github.com/bytebase/bytebase/plugin/advisor/db"
	"github.com/bytebase/bytebase/plugin/db"
)

// newTaskCheckStatementAdvisorSimpleExecutor creates a task check statement simple advisor executor.
func newTaskCheckStatementAdvisorSimpleExecutor() taskCheckExecutor {
	return &taskCheckStatementAdvisorSimpleExecutor{}
}

// taskCheckStatementAdvisorSimpleExecutor is the task check statement advisor simple executor.
type taskCheckStatementAdvisorSimpleExecutor struct {
}

// Run will run the task check statement advisor executor once.
func (*taskCheckStatementAdvisorSimpleExecutor) Run(_ context.Context, taskCheckRun *api.TaskCheckRun) (result []api.TaskCheckResult, err error) {
	payload := &api.TaskCheckDatabaseStatementAdvisePayload{}
	if err := json.Unmarshal([]byte(taskCheckRun.Payload), payload); err != nil {
		return nil, common.Wrapf(err, common.Invalid, "invalid check statement advise payload")
	}

	var advisorType advisor.Type
	switch taskCheckRun.Type {
	case api.TaskCheckDatabaseStatementFakeAdvise:
		advisorType = advisor.Fake
	case api.TaskCheckDatabaseStatementSyntax:
		switch payload.DbType {
		case db.MySQL, db.TiDB:
			advisorType = advisor.MySQLSyntax
		case db.Postgres:
			advisorType = advisor.PostgreSQLSyntax
		default:
			return nil, common.Errorf(common.Invalid, "invalid database type: %s for syntax statement advisor", payload.DbType)
		}
	}

	dbType, err := advisorDB.ConvertToAdvisorDBType(string(payload.DbType))
	if err != nil {
		return nil, err
	}

	adviceList, err := advisor.Check(
		dbType,
		advisorType,
		advisor.Context{
			Charset:   payload.Charset,
			Collation: payload.Collation,
		},
		payload.Statement,
	)
	if err != nil {
		return nil, common.Wrapf(err, common.Internal, "failed to check statement")
	}

	result = []api.TaskCheckResult{}
	for _, advice := range adviceList {
		status := api.TaskCheckStatusSuccess
		switch advice.Status {
		case advisor.Success:
			status = api.TaskCheckStatusSuccess
		case advisor.Warn:
			status = api.TaskCheckStatusWarn
		case advisor.Error:
			status = api.TaskCheckStatusError
		}

		result = append(result, api.TaskCheckResult{
			Status:    status,
			Namespace: api.AdvisorNamespace,
			Code:      advice.Code.Int(),
			Title:     advice.Title,
			Content:   advice.Content,
		})
	}

	return result, nil
}