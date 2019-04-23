$(document).ready(function() {
  $('#main_department').on('change', function() {
    var selectedDepartment = $(this).children("option:selected").val();
    var secondaryAvailableDepartments = '<option value="" disabled selected>الجهه الفرعيه</option>';
    secondaryDepartments.forEach(function(department) {
      if (department.MainDepartmentID === parseInt(selectedDepartment)) {
        var id = department.ID;
        var name = department.Name

        secondaryAvailableDepartments += `<option value = "${id}">${name}</option>`
      }
    })
    $('#secondary_department').html(secondaryAvailableDepartments);
  });
});
