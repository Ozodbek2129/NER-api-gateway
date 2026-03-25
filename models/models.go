package models

import "mime/multipart"

type File struct {
	File multipart.FileHeader `form:"file" binding:"required"`
}

type CreateContract struct {
	ContractName      string                `form:"contract_name" binding:"required"`
	ContractNumber    string                `form:"contract_number" binding:"required"`
	ContractDeadline  string                `form:"contract_deadline" binding:"required"`
	ResponsiblePerson string                `form:"responsible_person" binding:"required"`
	File              *multipart.FileHeader `form:"file" binding:"required"`
}