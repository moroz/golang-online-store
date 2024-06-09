#!/usr/bin/env -S bash -eu

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

CSV_SOURCE="$(realpath $SCRIPTPATH/../db/menu.csv)"

[ -f $CSV_SOURCE ]
echo $CSV_SOURCE

psql $DATABASE_URL <<-SQL
select 2+2;
SQL
