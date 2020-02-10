package model

import "github.com/hanpama/pgmg/introspect"

func NewSchemaFromIntrospection(ins *introspect.Schema) *Schema {
	var schema = &Schema{
		Name:   ins.SchemaName,
		Tables: make([]*Table, len(ins.Tables)),
	}

	typemap := make(map[string]*introspect.Type)
	for _, t := range ins.Types {
		typemap[t.SQLType] = t
	}

	tcs := make(map[string][]*Column)
	for _, c := range ins.Columns {
		it := typemap[c.DataType]

		t := Type{
			BaseType:   it.Name,
			FilterType: it.NullableName,
			SQLType:    it.SQLType,
			module:     it.Module,
		}
		if c.Default != "" {
			t.SelectType = it.Name
			t.InsertType = it.NullableName
		} else if c.IsNullable {
			t.SelectType = it.NullableName
			t.InsertType = it.NullableName
		} else {
			t.SelectType = it.Name
			t.InsertType = it.Name
		}

		tcs[c.TableName] = append(tcs[c.TableName], &Column{
			SQLName:      c.ColumnName,
			SQLDefault:   c.Default,
			SQLNullable:  c.IsNullable,
			SQLUpdatable: c.IsUpdatable,
			Type:         t,
		})
	}

	ks := make(map[string][]*Key)

	for _, ik := range ins.Keys {
		k := &Key{
			SQLName:      ik.ConstraintName,
			IsPrimaryKey: ik.IsPrimaryKey,
		}
		for _, cn := range ik.ColumnNames {
			for _, c := range tcs[ik.TableName] {
				if c.SQLName == cn {
					k.Columns = append(k.Columns, c)
				}
			}
		}

		ks[ik.TableName] = append(ks[ik.TableName], k)
	}

	fks := make(map[string][]*ForeignKey)

	for _, ifk := range ins.ForeignKeys {

		fk := &ForeignKey{
			SQLName: ifk.ConstraintName,
		}
		for _, cn := range ifk.ColumnNames {
			for _, c := range tcs[ifk.TableName] {
				if c.SQLName == cn {
					fk.Columns = append(fk.Columns, c)
				}
			}
		}

		for _, k := range ks[ifk.TargetTableName] {
			if k.SQLName == ifk.TargetConstraintName {
				fk.TargetKey = k
			}
		}

		fks[ifk.TableName] = append(fks[ifk.TableName], fk)
	}

	for i, it := range ins.Tables {
		t := new(Table)
		t.schema = schema
		t.Keys = ks[it.TableName]

		for _, k := range t.Keys {
			k.Table = t
		}

		t.IsInsertable = it.Insertable
		t.Columns = tcs[it.TableName]
		t.SQLName = it.TableName
		t.ForeignKeys = fks[it.TableName]

		schema.Tables[i] = t
	}

	return schema
}
