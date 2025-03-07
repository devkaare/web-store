package user

import "fmt"

func (r *UserRepo) DeleteUserByUserID(userID uint32) error {
	result, err := r.Client.Exec("DELETE FROM users WHERE user_id = $1", userID)
	if err != nil {
		return fmt.Errorf("DeleteUserByUserID %d, %v", userID, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteUserByUserID %d: %v", userID, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteUserByUserID %d: user not found", userID)
	}
	return nil
}
