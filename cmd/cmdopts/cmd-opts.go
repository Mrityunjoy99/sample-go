package cmdopts

type CmdOptsEnum string

const (
	MigrateAllOpts  CmdOptsEnum = "migrate-all"
	MigrateToOpts   CmdOptsEnum = "migrate-to"
	MigrateLastOpts CmdOptsEnum = "rollback-last"
	RollbackToOpts  CmdOptsEnum = "rollback-to"
	CommsWorker     CmdOptsEnum = "comms-worker"
)

func (e CmdOptsEnum) ToString() string {
	switch e {
	case MigrateAllOpts:
		return string(MigrateAllOpts)
	case MigrateToOpts:
		return string(MigrateToOpts)
	case MigrateLastOpts:
		return string(MigrateLastOpts)
	case RollbackToOpts:
		return string(RollbackToOpts)
	case CommsWorker:
		return string(CommsWorker)
	default:
		panic("invalid CmdOptsEnum")
	}
}
