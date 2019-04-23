package main

func findAllSecondaryDepartments() [] SecondaryDepartment {
  var secondaryDepartments [] SecondaryDepartment
  db.Find(&secondaryDepartments)
  return secondaryDepartments
}
