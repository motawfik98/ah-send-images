package main

func findAllMainDepartments() [] MainDepartment {
  var mainDepartments [] MainDepartment
  db.Find(&mainDepartments)
  return mainDepartments
}
