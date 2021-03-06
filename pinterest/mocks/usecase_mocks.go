// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/usecase_model.go

// Package mocks is a generated GoMock package.
package mocks

import (
	json "encoding/json"
	web_socket "github.com/go-park-mail-ru/2019_2_Solar/pinterest/web_socket"
	models "github.com/go-park-mail-ru/2019_2_Solar/pkg/models"
	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
	io "io"
	http "net/http"
	reflect "reflect"
)

// MockUseInterface is a mock of UseInterface interface
type MockUseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUseInterfaceMockRecorder
}

// MockUseInterfaceMockRecorder is the mock recorder for MockUseInterface
type MockUseInterfaceMockRecorder struct {
	mock *MockUseInterface
}

// NewMockUseInterface creates a new mock instance
func NewMockUseInterface(ctrl *gomock.Controller) *MockUseInterface {
	mock := &MockUseInterface{ctrl: ctrl}
	mock.recorder = &MockUseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseInterface) EXPECT() *MockUseInterfaceMockRecorder {
	return m.recorder
}

// SetJSONData mocks base method
func (m *MockUseInterface) SetJSONData(data interface{}, token, infMsg string) models.OutJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetJSONData", data, token, infMsg)
	ret0, _ := ret[0].(models.OutJSON)
	return ret0
}

// SetJSONData indicates an expected call of SetJSONData
func (mr *MockUseInterfaceMockRecorder) SetJSONData(data, token, infMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJSONData", reflect.TypeOf((*MockUseInterface)(nil).SetJSONData), data, token, infMsg)
}

// SetResponseError mocks base method
func (m *MockUseInterface) SetResponseError(encoder *json.Encoder, msg string, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetResponseError", encoder, msg, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetResponseError indicates an expected call of SetResponseError
func (mr *MockUseInterfaceMockRecorder) SetResponseError(encoder, msg, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResponseError", reflect.TypeOf((*MockUseInterface)(nil).SetResponseError), encoder, msg, err)
}

// GetUserByUsername mocks base method
func (m *MockUseInterface) GetUserByUsername(username string) (models.AnotherUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", username)
	ret0, _ := ret[0].(models.AnotherUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername
func (mr *MockUseInterfaceMockRecorder) GetUserByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUseInterface)(nil).GetUserByUsername), username)
}

// GetUserByEmail mocks base method
func (m *MockUseInterface) GetUserByEmail(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockUseInterfaceMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUseInterface)(nil).GetUserByEmail), email)
}

// GetUserIDByEmail mocks base method
func (m *MockUseInterface) GetUserIDByEmail(email string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDByEmail", email)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDByEmail indicates an expected call of GetUserIDByEmail
func (mr *MockUseInterfaceMockRecorder) GetUserIDByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDByEmail", reflect.TypeOf((*MockUseInterface)(nil).GetUserIDByEmail), email)
}

// GetAllUsers mocks base method
func (m *MockUseInterface) GetAllUsers() ([]models.AnotherUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]models.AnotherUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockUseInterfaceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUseInterface)(nil).GetAllUsers))
}

// ComparePassword mocks base method
func (m *MockUseInterface) ComparePassword(password, salt, loginPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePassword", password, salt, loginPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComparePassword indicates an expected call of ComparePassword
func (mr *MockUseInterfaceMockRecorder) ComparePassword(password, salt, loginPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*MockUseInterface)(nil).ComparePassword), password, salt, loginPassword)
}

// CheckRegDataValidation mocks base method
func (m *MockUseInterface) CheckRegDataValidation(newUser *models.UserReg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRegDataValidation", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRegDataValidation indicates an expected call of CheckRegDataValidation
func (mr *MockUseInterfaceMockRecorder) CheckRegDataValidation(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRegDataValidation", reflect.TypeOf((*MockUseInterface)(nil).CheckRegDataValidation), newUser)
}

// CheckRegUsernameEmailIsUnique mocks base method
func (m *MockUseInterface) CheckRegUsernameEmailIsUnique(username, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRegUsernameEmailIsUnique", username, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRegUsernameEmailIsUnique indicates an expected call of CheckRegUsernameEmailIsUnique
func (mr *MockUseInterfaceMockRecorder) CheckRegUsernameEmailIsUnique(username, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRegUsernameEmailIsUnique", reflect.TypeOf((*MockUseInterface)(nil).CheckRegUsernameEmailIsUnique), username, email)
}

// CheckProfileData mocks base method
func (m *MockUseInterface) CheckProfileData(newProfileUser *models.EditUserProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckProfileData", newProfileUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckProfileData indicates an expected call of CheckProfileData
func (mr *MockUseInterfaceMockRecorder) CheckProfileData(newProfileUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckProfileData", reflect.TypeOf((*MockUseInterface)(nil).CheckProfileData), newProfileUser)
}

// CheckUsernameEmailIsUnique mocks base method
func (m *MockUseInterface) CheckUsernameEmailIsUnique(newUsername, newEmail, username, email string, userID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsernameEmailIsUnique", newUsername, newEmail, username, email, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUsernameEmailIsUnique indicates an expected call of CheckUsernameEmailIsUnique
func (mr *MockUseInterfaceMockRecorder) CheckUsernameEmailIsUnique(newUsername, newEmail, username, email, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsernameEmailIsUnique", reflect.TypeOf((*MockUseInterface)(nil).CheckUsernameEmailIsUnique), newUsername, newEmail, username, email, userID)
}

// CheckBoardData mocks base method
func (m *MockUseInterface) CheckBoardData(newBoard models.NewBoard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBoardData", newBoard)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckBoardData indicates an expected call of CheckBoardData
func (mr *MockUseInterfaceMockRecorder) CheckBoardData(newBoard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBoardData", reflect.TypeOf((*MockUseInterface)(nil).CheckBoardData), newBoard)
}

// CheckPinData mocks base method
func (m *MockUseInterface) CheckPinData(newPin models.NewPin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPinData", newPin)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPinData indicates an expected call of CheckPinData
func (mr *MockUseInterfaceMockRecorder) CheckPinData(newPin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPinData", reflect.TypeOf((*MockUseInterface)(nil).CheckPinData), newPin)
}

// SetUserAvatarDir mocks base method
func (m *MockUseInterface) SetUserAvatarDir(idUser uint64, fileName string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserAvatarDir", idUser, fileName)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUserAvatarDir indicates an expected call of SetUserAvatarDir
func (mr *MockUseInterfaceMockRecorder) SetUserAvatarDir(idUser, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAvatarDir", reflect.TypeOf((*MockUseInterface)(nil).SetUserAvatarDir), idUser, fileName)
}

// SetUser mocks base method
func (m *MockUseInterface) SetUser(newUser models.EditUserProfile, user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUser", newUser, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUser indicates an expected call of SetUser
func (mr *MockUseInterfaceMockRecorder) SetUser(newUser, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUser", reflect.TypeOf((*MockUseInterface)(nil).SetUser), newUser, user)
}

// AddNewUser mocks base method
func (m *MockUseInterface) AddNewUser(username, email, password string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewUser", username, email, password)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewUser indicates an expected call of AddNewUser
func (mr *MockUseInterfaceMockRecorder) AddNewUser(username, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewUser", reflect.TypeOf((*MockUseInterface)(nil).AddNewUser), username, email, password)
}

// AddNewUserSession mocks base method
func (m *MockUseInterface) AddNewUserSession(userID uint64) (http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewUserSession", userID)
	ret0, _ := ret[0].(http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewUserSession indicates an expected call of AddNewUserSession
func (mr *MockUseInterfaceMockRecorder) AddNewUserSession(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewUserSession", reflect.TypeOf((*MockUseInterface)(nil).AddNewUserSession), userID)
}

// AddBoard mocks base method
func (m *MockUseInterface) AddBoard(newBoard models.Board) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBoard", newBoard)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBoard indicates an expected call of AddBoard
func (mr *MockUseInterfaceMockRecorder) AddBoard(newBoard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBoard", reflect.TypeOf((*MockUseInterface)(nil).AddBoard), newBoard)
}

// GetBoard mocks base method
func (m *MockUseInterface) GetBoard(boardID uint64) (models.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoard", boardID)
	ret0, _ := ret[0].(models.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoard indicates an expected call of GetBoard
func (mr *MockUseInterfaceMockRecorder) GetBoard(boardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoard", reflect.TypeOf((*MockUseInterface)(nil).GetBoard), boardID)
}

// GetMyBoards mocks base method
func (m *MockUseInterface) GetMyBoards(UserID uint64) ([]models.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyBoards", UserID)
	ret0, _ := ret[0].([]models.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyBoards indicates an expected call of GetMyBoards
func (mr *MockUseInterfaceMockRecorder) GetMyBoards(UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyBoards", reflect.TypeOf((*MockUseInterface)(nil).GetMyBoards), UserID)
}

// AddPin mocks base method
func (m *MockUseInterface) AddPin(newPin models.Pin) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPin", newPin)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPin indicates an expected call of AddPin
func (mr *MockUseInterfaceMockRecorder) AddPin(newPin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPin", reflect.TypeOf((*MockUseInterface)(nil).AddPin), newPin)
}

// GetPin mocks base method
func (m *MockUseInterface) GetPin(pinID uint64) (models.FullPin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPin", pinID)
	ret0, _ := ret[0].(models.FullPin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPin indicates an expected call of GetPin
func (mr *MockUseInterfaceMockRecorder) GetPin(pinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPin", reflect.TypeOf((*MockUseInterface)(nil).GetPin), pinID)
}

// GetPinsDisplay mocks base method
func (m *MockUseInterface) GetPinsDisplay(boardID uint64) ([]models.PinDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPinsDisplay", boardID)
	ret0, _ := ret[0].([]models.PinDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinsDisplay indicates an expected call of GetPinsDisplay
func (mr *MockUseInterfaceMockRecorder) GetPinsDisplay(boardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinsDisplay", reflect.TypeOf((*MockUseInterface)(nil).GetPinsDisplay), boardID)
}

// GetPinsByUsername mocks base method
func (m *MockUseInterface) GetPinsByUsername(useID int) ([]models.PinDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPinsByUsername", useID)
	ret0, _ := ret[0].([]models.PinDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinsByUsername indicates an expected call of GetPinsByUsername
func (mr *MockUseInterfaceMockRecorder) GetPinsByUsername(useID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinsByUsername", reflect.TypeOf((*MockUseInterface)(nil).GetPinsByUsername), useID)
}

// GetNewPins mocks base method
func (m *MockUseInterface) GetNewPins() ([]models.PinDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewPins")
	ret0, _ := ret[0].([]models.PinDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewPins indicates an expected call of GetNewPins
func (mr *MockUseInterfaceMockRecorder) GetNewPins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewPins", reflect.TypeOf((*MockUseInterface)(nil).GetNewPins))
}

// GetMyPins mocks base method
func (m *MockUseInterface) GetMyPins(userID uint64) ([]models.PinDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyPins", userID)
	ret0, _ := ret[0].([]models.PinDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyPins indicates an expected call of GetMyPins
func (mr *MockUseInterfaceMockRecorder) GetMyPins(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyPins", reflect.TypeOf((*MockUseInterface)(nil).GetMyPins), userID)
}

// GetSubscribePins mocks base method
func (m *MockUseInterface) GetSubscribePins(userID uint64) ([]models.PinDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribePins", userID)
	ret0, _ := ret[0].([]models.PinDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribePins indicates an expected call of GetSubscribePins
func (mr *MockUseInterfaceMockRecorder) GetSubscribePins(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribePins", reflect.TypeOf((*MockUseInterface)(nil).GetSubscribePins), userID)
}

// AddComment mocks base method
func (m *MockUseInterface) AddComment(pinID, userID uint64, newComment models.NewComment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", pinID, userID, newComment)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddComment indicates an expected call of AddComment
func (mr *MockUseInterfaceMockRecorder) AddComment(pinID, userID, newComment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockUseInterface)(nil).AddComment), pinID, userID, newComment)
}

// GetComments mocks base method
func (m *MockUseInterface) GetComments(pinID uint64) ([]models.CommentDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", pinID)
	ret0, _ := ret[0].([]models.CommentDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments
func (mr *MockUseInterfaceMockRecorder) GetComments(pinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockUseInterface)(nil).GetComments), pinID)
}

// AddNotice mocks base method
func (m *MockUseInterface) AddNotice(newNotice models.Notice) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNotice", newNotice)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNotice indicates an expected call of AddNotice
func (mr *MockUseInterfaceMockRecorder) AddNotice(newNotice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNotice", reflect.TypeOf((*MockUseInterface)(nil).AddNotice), newNotice)
}

// GetMyNotices mocks base method
func (m *MockUseInterface) GetMyNotices(userID uint64) ([]models.Notice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyNotices", userID)
	ret0, _ := ret[0].([]models.Notice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyNotices indicates an expected call of GetMyNotices
func (mr *MockUseInterfaceMockRecorder) GetMyNotices(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyNotices", reflect.TypeOf((*MockUseInterface)(nil).GetMyNotices), userID)
}

// AddSubscribe mocks base method
func (m *MockUseInterface) AddSubscribe(userID uint64, followeeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscribe", userID, followeeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSubscribe indicates an expected call of AddSubscribe
func (mr *MockUseInterfaceMockRecorder) AddSubscribe(userID, followeeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscribe", reflect.TypeOf((*MockUseInterface)(nil).AddSubscribe), userID, followeeName)
}

// RemoveSubscribe mocks base method
func (m *MockUseInterface) RemoveSubscribe(userID uint64, followeeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSubscribe", userID, followeeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSubscribe indicates an expected call of RemoveSubscribe
func (mr *MockUseInterfaceMockRecorder) RemoveSubscribe(userID, followeeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubscribe", reflect.TypeOf((*MockUseInterface)(nil).RemoveSubscribe), userID, followeeName)
}

// ExtractFormatFile mocks base method
func (m *MockUseInterface) ExtractFormatFile(fileName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractFormatFile", fileName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractFormatFile indicates an expected call of ExtractFormatFile
func (mr *MockUseInterfaceMockRecorder) ExtractFormatFile(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractFormatFile", reflect.TypeOf((*MockUseInterface)(nil).ExtractFormatFile), fileName)
}

// RemoveOldUserSession mocks base method
func (m *MockUseInterface) RemoveOldUserSession(sessionKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOldUserSession", sessionKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveOldUserSession indicates an expected call of RemoveOldUserSession
func (mr *MockUseInterfaceMockRecorder) RemoveOldUserSession(sessionKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOldUserSession", reflect.TypeOf((*MockUseInterface)(nil).RemoveOldUserSession), sessionKey)
}

// CalculateMD5FromFile mocks base method
func (m *MockUseInterface) CalculateMD5FromFile(fileByte io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateMD5FromFile", fileByte)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateMD5FromFile indicates an expected call of CalculateMD5FromFile
func (mr *MockUseInterfaceMockRecorder) CalculateMD5FromFile(fileByte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateMD5FromFile", reflect.TypeOf((*MockUseInterface)(nil).CalculateMD5FromFile), fileByte)
}

// AddDir mocks base method
func (m *MockUseInterface) AddDir(folder string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDir", folder)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDir indicates an expected call of AddDir
func (mr *MockUseInterfaceMockRecorder) AddDir(folder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDir", reflect.TypeOf((*MockUseInterface)(nil).AddDir), folder)
}

// AddPictureFile mocks base method
func (m *MockUseInterface) AddPictureFile(fileName string, fileByte io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPictureFile", fileName, fileByte)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPictureFile indicates an expected call of AddPictureFile
func (mr *MockUseInterfaceMockRecorder) AddPictureFile(fileName, fileByte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPictureFile", reflect.TypeOf((*MockUseInterface)(nil).AddPictureFile), fileName, fileByte)
}

// ReturnHub mocks base method
func (m *MockUseInterface) ReturnHub() *web_socket.HubStruct {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnHub")
	ret0, _ := ret[0].(*web_socket.HubStruct)
	return ret0
}

// ReturnHub indicates an expected call of ReturnHub
func (mr *MockUseInterfaceMockRecorder) ReturnHub() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnHub", reflect.TypeOf((*MockUseInterface)(nil).ReturnHub))
}

// SearchPinsByTag mocks base method
func (m *MockUseInterface) SearchPinsByTag(tag string) ([]models.PinDisplay, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPinsByTag", tag)
	ret0, _ := ret[0].([]models.PinDisplay)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPinsByTag indicates an expected call of SearchPinsByTag
func (mr *MockUseInterfaceMockRecorder) SearchPinsByTag(tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPinsByTag", reflect.TypeOf((*MockUseInterface)(nil).SearchPinsByTag), tag)
}

// SearchUserByUsername mocks base method
func (m *MockUseInterface) SearchUserByUsername(username string) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUserByUsername", username)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUserByUsername indicates an expected call of SearchUserByUsername
func (mr *MockUseInterfaceMockRecorder) SearchUserByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUserByUsername", reflect.TypeOf((*MockUseInterface)(nil).SearchUserByUsername), username)
}

// CreateClient mocks base method
func (m *MockUseInterface) CreateClient(conn *websocket.Conn, user models.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateClient", conn, user)
}

// CreateClient indicates an expected call of CreateClient
func (mr *MockUseInterfaceMockRecorder) CreateClient(conn, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockUseInterface)(nil).CreateClient), conn, user)
}

// GetMySubscribeByUsername mocks base method
func (m *MockUseInterface) GetMySubscribeByUsername(userId uint64, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMySubscribeByUsername", userId, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMySubscribeByUsername indicates an expected call of GetMySubscribeByUsername
func (mr *MockUseInterfaceMockRecorder) GetMySubscribeByUsername(userId, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMySubscribeByUsername", reflect.TypeOf((*MockUseInterface)(nil).GetMySubscribeByUsername), userId, username)
}

// AddTags mocks base method
func (m *MockUseInterface) AddTags(description string, pinID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTags", description, pinID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTags indicates an expected call of AddTags
func (mr *MockUseInterfaceMockRecorder) AddTags(description, pinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTags", reflect.TypeOf((*MockUseInterface)(nil).AddTags), description, pinID)
}

// GetCategories mocks base method
func (m *MockUseInterface) GetCategories() ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories
func (mr *MockUseInterfaceMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockUseInterface)(nil).GetCategories))
}

// GetMessages mocks base method
func (m *MockUseInterface) GetMessages(senderId, receiverId uint64) ([]models.OutputMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", senderId, receiverId)
	ret0, _ := ret[0].([]models.OutputMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages
func (mr *MockUseInterfaceMockRecorder) GetMessages(senderId, receiverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockUseInterface)(nil).GetMessages), senderId, receiverId)
}

// GetUserByCookieValue mocks base method
func (m *MockUseInterface) GetUserByCookieValue(cookieValue string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByCookieValue", cookieValue)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByCookieValue indicates an expected call of GetUserByCookieValue
func (mr *MockUseInterfaceMockRecorder) GetUserByCookieValue(cookieValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByCookieValue", reflect.TypeOf((*MockUseInterface)(nil).GetUserByCookieValue), cookieValue)
}

// GetSessionsByCookieValue mocks base method
func (m *MockUseInterface) GetSessionsByCookieValue(cookieValue string) (models.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionsByCookieValue", cookieValue)
	ret0, _ := ret[0].(models.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionsByCookieValue indicates an expected call of GetSessionsByCookieValue
func (mr *MockUseInterfaceMockRecorder) GetSessionsByCookieValue(cookieValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionsByCookieValue", reflect.TypeOf((*MockUseInterface)(nil).GetSessionsByCookieValue), cookieValue)
}

// MGetSessionsByCookieValue mocks base method
func (m *MockUseInterface) MGetSessionsByCookieValue(cookieValue string) (models.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MGetSessionsByCookieValue", cookieValue)
	ret0, _ := ret[0].(models.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MGetSessionsByCookieValue indicates an expected call of MGetSessionsByCookieValue
func (mr *MockUseInterfaceMockRecorder) MGetSessionsByCookieValue(cookieValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGetSessionsByCookieValue", reflect.TypeOf((*MockUseInterface)(nil).MGetSessionsByCookieValue), cookieValue)
}
