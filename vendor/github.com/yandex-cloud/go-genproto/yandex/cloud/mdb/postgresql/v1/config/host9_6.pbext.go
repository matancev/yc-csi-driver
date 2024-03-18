// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlHostConfig9_6) SetRecoveryMinApplyDelay(v *wrapperspb.Int64Value) {
	m.RecoveryMinApplyDelay = v
}

func (m *PostgresqlHostConfig9_6) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlHostConfig9_6) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlHostConfig9_6) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlHostConfig9_6) SetReplacementSortTuples(v *wrapperspb.Int64Value) {
	m.ReplacementSortTuples = v
}

func (m *PostgresqlHostConfig9_6) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlHostConfig9_6) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlHostConfig9_6) SetOldSnapshotThreshold(v *wrapperspb.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlHostConfig9_6) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlHostConfig9_6) SetConstraintExclusion(v PostgresqlHostConfig9_6_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlHostConfig9_6) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlHostConfig9_6) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlHostConfig9_6) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlHostConfig9_6) SetForceParallelMode(v PostgresqlHostConfig9_6_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlHostConfig9_6) SetClientMinMessages(v PostgresqlHostConfig9_6_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlHostConfig9_6) SetLogMinMessages(v PostgresqlHostConfig9_6_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlHostConfig9_6) SetLogMinErrorStatement(v PostgresqlHostConfig9_6_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlHostConfig9_6) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlHostConfig9_6) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlHostConfig9_6) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlHostConfig9_6) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlHostConfig9_6) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlHostConfig9_6) SetLogErrorVerbosity(v PostgresqlHostConfig9_6_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlHostConfig9_6) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlHostConfig9_6) SetLogStatement(v PostgresqlHostConfig9_6_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlHostConfig9_6) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlHostConfig9_6) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlHostConfig9_6) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlHostConfig9_6) SetDefaultTransactionIsolation(v PostgresqlHostConfig9_6_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlHostConfig9_6) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlHostConfig9_6) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlHostConfig9_6) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlHostConfig9_6) SetByteaOutput(v PostgresqlHostConfig9_6_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlHostConfig9_6) SetXmlbinary(v PostgresqlHostConfig9_6_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlHostConfig9_6) SetXmloption(v PostgresqlHostConfig9_6_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlHostConfig9_6) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlHostConfig9_6) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlHostConfig9_6) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlHostConfig9_6) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlHostConfig9_6) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlHostConfig9_6) SetBackslashQuote(v PostgresqlHostConfig9_6_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlHostConfig9_6) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlHostConfig9_6) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlHostConfig9_6) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlHostConfig9_6) SetOperatorPrecedenceWarning(v *wrapperspb.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlHostConfig9_6) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlHostConfig9_6) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlHostConfig9_6) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlHostConfig9_6) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlHostConfig9_6) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlHostConfig9_6) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlHostConfig9_6) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlHostConfig9_6) SetSqlInheritance(v *wrapperspb.BoolValue) {
	m.SqlInheritance = v
}

func (m *PostgresqlHostConfig9_6) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlHostConfig9_6) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}