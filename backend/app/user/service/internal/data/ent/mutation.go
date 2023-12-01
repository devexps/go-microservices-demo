// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/devexps/go-microservices-demo/app/user/service/internal/data/ent/predicate"
	"github.com/devexps/go-microservices-demo/app/user/service/internal/data/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uint32
	create_by     *uint32
	addcreate_by  *int32
	create_time   *time.Time
	update_time   *time.Time
	delete_time   *time.Time
	status        *user.Status
	username      *string
	password      *string
	nick_name     *string
	real_name     *string
	email         *string
	phone         *string
	authority     *user.Authority
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uint32) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uint32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uint32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uint32, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint32{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateBy sets the "create_by" field.
func (m *UserMutation) SetCreateBy(u uint32) {
	m.create_by = &u
	m.addcreate_by = nil
}

// CreateBy returns the value of the "create_by" field in the mutation.
func (m *UserMutation) CreateBy() (r uint32, exists bool) {
	v := m.create_by
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateBy returns the old "create_by" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreateBy(ctx context.Context) (v *uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateBy is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateBy requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateBy: %w", err)
	}
	return oldValue.CreateBy, nil
}

// AddCreateBy adds u to the "create_by" field.
func (m *UserMutation) AddCreateBy(u int32) {
	if m.addcreate_by != nil {
		*m.addcreate_by += u
	} else {
		m.addcreate_by = &u
	}
}

// AddedCreateBy returns the value that was added to the "create_by" field in this mutation.
func (m *UserMutation) AddedCreateBy() (r int32, exists bool) {
	v := m.addcreate_by
	if v == nil {
		return
	}
	return *v, true
}

// ClearCreateBy clears the value of the "create_by" field.
func (m *UserMutation) ClearCreateBy() {
	m.create_by = nil
	m.addcreate_by = nil
	m.clearedFields[user.FieldCreateBy] = struct{}{}
}

// CreateByCleared returns if the "create_by" field was cleared in this mutation.
func (m *UserMutation) CreateByCleared() bool {
	_, ok := m.clearedFields[user.FieldCreateBy]
	return ok
}

// ResetCreateBy resets all changes to the "create_by" field.
func (m *UserMutation) ResetCreateBy() {
	m.create_by = nil
	m.addcreate_by = nil
	delete(m.clearedFields, user.FieldCreateBy)
}

// SetCreateTime sets the "create_time" field.
func (m *UserMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *UserMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreateTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ClearCreateTime clears the value of the "create_time" field.
func (m *UserMutation) ClearCreateTime() {
	m.create_time = nil
	m.clearedFields[user.FieldCreateTime] = struct{}{}
}

// CreateTimeCleared returns if the "create_time" field was cleared in this mutation.
func (m *UserMutation) CreateTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldCreateTime]
	return ok
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *UserMutation) ResetCreateTime() {
	m.create_time = nil
	delete(m.clearedFields, user.FieldCreateTime)
}

// SetUpdateTime sets the "update_time" field.
func (m *UserMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *UserMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdateTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ClearUpdateTime clears the value of the "update_time" field.
func (m *UserMutation) ClearUpdateTime() {
	m.update_time = nil
	m.clearedFields[user.FieldUpdateTime] = struct{}{}
}

// UpdateTimeCleared returns if the "update_time" field was cleared in this mutation.
func (m *UserMutation) UpdateTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldUpdateTime]
	return ok
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *UserMutation) ResetUpdateTime() {
	m.update_time = nil
	delete(m.clearedFields, user.FieldUpdateTime)
}

// SetDeleteTime sets the "delete_time" field.
func (m *UserMutation) SetDeleteTime(t time.Time) {
	m.delete_time = &t
}

// DeleteTime returns the value of the "delete_time" field in the mutation.
func (m *UserMutation) DeleteTime() (r time.Time, exists bool) {
	v := m.delete_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteTime returns the old "delete_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDeleteTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleteTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleteTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteTime: %w", err)
	}
	return oldValue.DeleteTime, nil
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (m *UserMutation) ClearDeleteTime() {
	m.delete_time = nil
	m.clearedFields[user.FieldDeleteTime] = struct{}{}
}

// DeleteTimeCleared returns if the "delete_time" field was cleared in this mutation.
func (m *UserMutation) DeleteTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldDeleteTime]
	return ok
}

// ResetDeleteTime resets all changes to the "delete_time" field.
func (m *UserMutation) ResetDeleteTime() {
	m.delete_time = nil
	delete(m.clearedFields, user.FieldDeleteTime)
}

// SetStatus sets the "status" field.
func (m *UserMutation) SetStatus(u user.Status) {
	m.status = &u
}

// Status returns the value of the "status" field in the mutation.
func (m *UserMutation) Status() (r user.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStatus(ctx context.Context) (v *user.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ClearStatus clears the value of the "status" field.
func (m *UserMutation) ClearStatus() {
	m.status = nil
	m.clearedFields[user.FieldStatus] = struct{}{}
}

// StatusCleared returns if the "status" field was cleared in this mutation.
func (m *UserMutation) StatusCleared() bool {
	_, ok := m.clearedFields[user.FieldStatus]
	return ok
}

// ResetStatus resets all changes to the "status" field.
func (m *UserMutation) ResetStatus() {
	m.status = nil
	delete(m.clearedFields, user.FieldStatus)
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ClearUsername clears the value of the "username" field.
func (m *UserMutation) ClearUsername() {
	m.username = nil
	m.clearedFields[user.FieldUsername] = struct{}{}
}

// UsernameCleared returns if the "username" field was cleared in this mutation.
func (m *UserMutation) UsernameCleared() bool {
	_, ok := m.clearedFields[user.FieldUsername]
	return ok
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
	delete(m.clearedFields, user.FieldUsername)
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ClearPassword clears the value of the "password" field.
func (m *UserMutation) ClearPassword() {
	m.password = nil
	m.clearedFields[user.FieldPassword] = struct{}{}
}

// PasswordCleared returns if the "password" field was cleared in this mutation.
func (m *UserMutation) PasswordCleared() bool {
	_, ok := m.clearedFields[user.FieldPassword]
	return ok
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
	delete(m.clearedFields, user.FieldPassword)
}

// SetNickName sets the "nick_name" field.
func (m *UserMutation) SetNickName(s string) {
	m.nick_name = &s
}

// NickName returns the value of the "nick_name" field in the mutation.
func (m *UserMutation) NickName() (r string, exists bool) {
	v := m.nick_name
	if v == nil {
		return
	}
	return *v, true
}

// OldNickName returns the old "nick_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickName: %w", err)
	}
	return oldValue.NickName, nil
}

// ClearNickName clears the value of the "nick_name" field.
func (m *UserMutation) ClearNickName() {
	m.nick_name = nil
	m.clearedFields[user.FieldNickName] = struct{}{}
}

// NickNameCleared returns if the "nick_name" field was cleared in this mutation.
func (m *UserMutation) NickNameCleared() bool {
	_, ok := m.clearedFields[user.FieldNickName]
	return ok
}

// ResetNickName resets all changes to the "nick_name" field.
func (m *UserMutation) ResetNickName() {
	m.nick_name = nil
	delete(m.clearedFields, user.FieldNickName)
}

// SetRealName sets the "real_name" field.
func (m *UserMutation) SetRealName(s string) {
	m.real_name = &s
}

// RealName returns the value of the "real_name" field in the mutation.
func (m *UserMutation) RealName() (r string, exists bool) {
	v := m.real_name
	if v == nil {
		return
	}
	return *v, true
}

// OldRealName returns the old "real_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRealName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRealName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRealName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRealName: %w", err)
	}
	return oldValue.RealName, nil
}

// ClearRealName clears the value of the "real_name" field.
func (m *UserMutation) ClearRealName() {
	m.real_name = nil
	m.clearedFields[user.FieldRealName] = struct{}{}
}

// RealNameCleared returns if the "real_name" field was cleared in this mutation.
func (m *UserMutation) RealNameCleared() bool {
	_, ok := m.clearedFields[user.FieldRealName]
	return ok
}

// ResetRealName resets all changes to the "real_name" field.
func (m *UserMutation) ResetRealName() {
	m.real_name = nil
	delete(m.clearedFields, user.FieldRealName)
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ClearEmail clears the value of the "email" field.
func (m *UserMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[user.FieldEmail] = struct{}{}
}

// EmailCleared returns if the "email" field was cleared in this mutation.
func (m *UserMutation) EmailCleared() bool {
	_, ok := m.clearedFields[user.FieldEmail]
	return ok
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, user.FieldEmail)
}

// SetPhone sets the "phone" field.
func (m *UserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *UserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPhone(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ClearPhone clears the value of the "phone" field.
func (m *UserMutation) ClearPhone() {
	m.phone = nil
	m.clearedFields[user.FieldPhone] = struct{}{}
}

// PhoneCleared returns if the "phone" field was cleared in this mutation.
func (m *UserMutation) PhoneCleared() bool {
	_, ok := m.clearedFields[user.FieldPhone]
	return ok
}

// ResetPhone resets all changes to the "phone" field.
func (m *UserMutation) ResetPhone() {
	m.phone = nil
	delete(m.clearedFields, user.FieldPhone)
}

// SetAuthority sets the "authority" field.
func (m *UserMutation) SetAuthority(u user.Authority) {
	m.authority = &u
}

// Authority returns the value of the "authority" field in the mutation.
func (m *UserMutation) Authority() (r user.Authority, exists bool) {
	v := m.authority
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthority returns the old "authority" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAuthority(ctx context.Context) (v *user.Authority, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAuthority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAuthority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthority: %w", err)
	}
	return oldValue.Authority, nil
}

// ClearAuthority clears the value of the "authority" field.
func (m *UserMutation) ClearAuthority() {
	m.authority = nil
	m.clearedFields[user.FieldAuthority] = struct{}{}
}

// AuthorityCleared returns if the "authority" field was cleared in this mutation.
func (m *UserMutation) AuthorityCleared() bool {
	_, ok := m.clearedFields[user.FieldAuthority]
	return ok
}

// ResetAuthority resets all changes to the "authority" field.
func (m *UserMutation) ResetAuthority() {
	m.authority = nil
	delete(m.clearedFields, user.FieldAuthority)
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 12)
	if m.create_by != nil {
		fields = append(fields, user.FieldCreateBy)
	}
	if m.create_time != nil {
		fields = append(fields, user.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, user.FieldUpdateTime)
	}
	if m.delete_time != nil {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.status != nil {
		fields = append(fields, user.FieldStatus)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.nick_name != nil {
		fields = append(fields, user.FieldNickName)
	}
	if m.real_name != nil {
		fields = append(fields, user.FieldRealName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.phone != nil {
		fields = append(fields, user.FieldPhone)
	}
	if m.authority != nil {
		fields = append(fields, user.FieldAuthority)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreateBy:
		return m.CreateBy()
	case user.FieldCreateTime:
		return m.CreateTime()
	case user.FieldUpdateTime:
		return m.UpdateTime()
	case user.FieldDeleteTime:
		return m.DeleteTime()
	case user.FieldStatus:
		return m.Status()
	case user.FieldUsername:
		return m.Username()
	case user.FieldPassword:
		return m.Password()
	case user.FieldNickName:
		return m.NickName()
	case user.FieldRealName:
		return m.RealName()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPhone:
		return m.Phone()
	case user.FieldAuthority:
		return m.Authority()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreateBy:
		return m.OldCreateBy(ctx)
	case user.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case user.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case user.FieldDeleteTime:
		return m.OldDeleteTime(ctx)
	case user.FieldStatus:
		return m.OldStatus(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldNickName:
		return m.OldNickName(ctx)
	case user.FieldRealName:
		return m.OldRealName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPhone:
		return m.OldPhone(ctx)
	case user.FieldAuthority:
		return m.OldAuthority(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreateBy:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateBy(v)
		return nil
	case user.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case user.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case user.FieldDeleteTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteTime(v)
		return nil
	case user.FieldStatus:
		v, ok := value.(user.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldNickName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickName(v)
		return nil
	case user.FieldRealName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRealName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case user.FieldAuthority:
		v, ok := value.(user.Authority)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthority(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addcreate_by != nil {
		fields = append(fields, user.FieldCreateBy)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreateBy:
		return m.AddedCreateBy()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreateBy:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreateBy(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldCreateBy) {
		fields = append(fields, user.FieldCreateBy)
	}
	if m.FieldCleared(user.FieldCreateTime) {
		fields = append(fields, user.FieldCreateTime)
	}
	if m.FieldCleared(user.FieldUpdateTime) {
		fields = append(fields, user.FieldUpdateTime)
	}
	if m.FieldCleared(user.FieldDeleteTime) {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.FieldCleared(user.FieldStatus) {
		fields = append(fields, user.FieldStatus)
	}
	if m.FieldCleared(user.FieldUsername) {
		fields = append(fields, user.FieldUsername)
	}
	if m.FieldCleared(user.FieldPassword) {
		fields = append(fields, user.FieldPassword)
	}
	if m.FieldCleared(user.FieldNickName) {
		fields = append(fields, user.FieldNickName)
	}
	if m.FieldCleared(user.FieldRealName) {
		fields = append(fields, user.FieldRealName)
	}
	if m.FieldCleared(user.FieldEmail) {
		fields = append(fields, user.FieldEmail)
	}
	if m.FieldCleared(user.FieldPhone) {
		fields = append(fields, user.FieldPhone)
	}
	if m.FieldCleared(user.FieldAuthority) {
		fields = append(fields, user.FieldAuthority)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldCreateBy:
		m.ClearCreateBy()
		return nil
	case user.FieldCreateTime:
		m.ClearCreateTime()
		return nil
	case user.FieldUpdateTime:
		m.ClearUpdateTime()
		return nil
	case user.FieldDeleteTime:
		m.ClearDeleteTime()
		return nil
	case user.FieldStatus:
		m.ClearStatus()
		return nil
	case user.FieldUsername:
		m.ClearUsername()
		return nil
	case user.FieldPassword:
		m.ClearPassword()
		return nil
	case user.FieldNickName:
		m.ClearNickName()
		return nil
	case user.FieldRealName:
		m.ClearRealName()
		return nil
	case user.FieldEmail:
		m.ClearEmail()
		return nil
	case user.FieldPhone:
		m.ClearPhone()
		return nil
	case user.FieldAuthority:
		m.ClearAuthority()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreateBy:
		m.ResetCreateBy()
		return nil
	case user.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case user.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case user.FieldDeleteTime:
		m.ResetDeleteTime()
		return nil
	case user.FieldStatus:
		m.ResetStatus()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldNickName:
		m.ResetNickName()
		return nil
	case user.FieldRealName:
		m.ResetRealName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPhone:
		m.ResetPhone()
		return nil
	case user.FieldAuthority:
		m.ResetAuthority()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}