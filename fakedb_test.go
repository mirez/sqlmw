package sqlmw

import (
	"context"
	"database/sql/driver"
)

type fakeDriver struct {
	conn driver.Conn
}

func (d *fakeDriver) Open(_ string) (driver.Conn, error) {
	return d.conn, nil
}

type fakeStmt struct {
	called bool
}

type fakeStmtWithCheckNamedValue struct {
	fakeStmt
}

type fakeStmtWithoutCheckNamedValue struct {
	fakeStmt
}

func (s fakeStmt) Close() error {
	return nil
}

func (s fakeStmt) NumInput() int {
	return 1
}

func (s fakeStmt) Exec(_ []driver.Value) (driver.Result, error) {
	return nil, nil
}

func (s fakeStmt) Query(_ []driver.Value) (driver.Rows, error) {
	return nil, nil
}

func (s *fakeStmtWithCheckNamedValue) CheckNamedValue(_ *driver.NamedValue) (err error) {
	s.called = true
	return
}

type fakeConn struct {
	called bool
	stmt   driver.Stmt
}

type fakeConnWithCheckNamedValue struct {
	fakeConn
}

type fakeConnWithoutCheckNamedValue struct {
	fakeConn
}

func (c *fakeConn) Prepare(_ string) (driver.Stmt, error) {
	return nil, nil
}

func (c *fakeConn) PrepareContext(_ context.Context, _ string) (driver.Stmt, error) {
	return c.stmt, nil
}

func (c *fakeConn) Close() error              { return nil }

func (c *fakeConn) Begin() (driver.Tx, error) { return nil, nil }

func (c *fakeConn) QueryContext(_ context.Context, _ string, _ []driver.NamedValue) (driver.Rows, error) {
	return nil, nil
}

func (c *fakeConnWithCheckNamedValue) CheckNamedValue(_ *driver.NamedValue) (err error) {
	c.called = true
	return
}

type fakeInterceptor struct {
	NullInterceptor
}
