package fixtures

import "context"

// InitTestData init all test data
func InitTestData(ctx context.Context) error {
	err := AddUserTest(ctx)
	if err != nil {
		return err
	}

	err = AddTaskTest(ctx)
	if err != nil {
		return err
	}

	err = AddCommentTest(ctx)
	if err != nil {
		return err
	}

	return err
}
