$(document).ready(function () {
  $('#register').change(function () {
    if ($('#register').is(':checked')) {
      $('#nickname').attr('disabled', false);
      $('#animal').attr('disabled', false);
      $('#address').attr('disabled', false);
    } else {
      $('#nickname').attr('disabled', true);
      $('#animal').attr('disabled', true);
      $('#address').attr('disabled', true);
    }
  });
  $('#submit').click(function () {
    if ($('#register').is(':checked')) {
      let formData = new FormData();
      formData.append('phone', $('#phone').val());
      formData.append('nickname', $('#nickname').val());
      formData.append('animal', $('#animal').val());
      formData.append('address', $('#address').val());
      formData.append('password', $('#password').val());
      $.ajax({
        url: '/register',
        type: 'post',
        data: formData,
        processData: false,
        contentType: false,
        dataType: 'json',
        success: function (data) {
          if (data.pass) {
            window.localStorage.setItem("phone", data.phone)
            window.location.href = '/dash';
            alert('register successfully!');
          }else{
            console.log('register failed!');
          }
        },
        error: function () {
          console.log('register failed!');
        },
      });
    } else {
      let formData = new FormData();
      formData.append('phone', $('#phone').val());
      formData.append('password', $('#password').val());
      $.ajax({
        type: "POST",
        url: "/login",
        data: formData,
        processData: false,
        contentType: false,
        dataType: "json",
        success: function (data) {
          console.log('login successfully!');
          if (data.pass == true) {
            window.localStorage.setItem("phone", data.phone)
            window.location.href = '/dash';
          } else {
            alert("账号或密码错误")
            console.log('login failed!');
          }
        },
        error: function () {
          alert('submit failed!');
        },
      });
    }
  });
});
