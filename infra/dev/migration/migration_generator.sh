#!/bin/sh

echo "\nOutput of migration generator"

MIGRATION_NAME=$1
DATETIME=$(date +"%Y%m%d%H%M%S")

MIGRATION_VERSION="${DATETIME}_${MIGRATION_NAME}"
MIGRATION_FILE="./migration/${MIGRATION_VERSION}.go"

MIGRATION_GO_VAR="V${DATETIME}${MIGRATION_NAME}"

TEMPLATE="\
package migration\n\
\n\
import (\n\
\t\"github.com/go-gormigrate/gormigrate/v2\"\n\
\t\"gorm.io/gorm\"\n\
)\n\
\n\
var ${MIGRATION_GO_VAR} *gormigrate.Migration = &gormigrate.Migration{\n\
\tID: \"${MIGRATION_VERSION}\",\n\
\tMigrate: func(tx *gorm.DB) error {\n\
\t\treturn nil\n\
\t},\n\
\tRollback: func(tx *gorm.DB) error {\n\
\t\treturn nil\n\
\t},\n\
}\
"

echo "Scaffoliding migration"

echo $TEMPLATE > $MIGRATION_FILE

echo "migration successful"
