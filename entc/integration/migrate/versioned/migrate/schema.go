// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_car", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cars_users_car",
				Columns:    []*schema.Column{CarsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString, Size: 10},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "renamed", Type: field.TypeString, Nullable: true},
		{Name: "blob", Type: field.TypeBytes, Nullable: true, Size: 255},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"logged_in", "logged_out"}},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "workplace", Type: field.TypeString, Nullable: true, Size: 30},
		{Name: "user_children", Type: field.TypeInt, Nullable: true},
		{Name: "user_spouse", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_children",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_spouse",
				Columns:    []*schema.Column{UsersColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_description",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Prefix: 50,
				},
			},
			{
				Name:    "user_name_address",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2], UsersColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		UsersTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[1].RefTable = UsersTable
}