$(document).ready(function () {

  // $('.ui.form').form({
  //     fields: {
  //       email: 'empty',
  //       password: 'empty',
  //       password_confirm: 'empty'
  //     }
  //   });

  // $('.input').on('click', function () {
  //   $('.ui.error.message').empty().hide()
  // });

  // $("#submit").on('click', function () {

  //     var _email, _password
  //     $(this).addClass('loading');
  //     var $form = $('.ui.form')
  //     _email = $form.form('get value', 'email')
  //     _password = $form.form('get value', 'password')

  //   if (_email == '' || _password == '') {
  //     $('.ui.error.message').empty().append("Email and Passwor must have value").show()
  //   } else {

  //     var data = {
  //       email: _email,
  //       password: _password
  //     };

  //     var jqxhr = $.ajax({
  //       method: "POST",
  //       url: "/v1/jwt/form/login",
  //       data: data,
  //       // beforeSend: function (request) {
  //       //   request.setRequestHeader("Authoriztion", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTA0NjI4MjI5LCJpYXQiOjE1MDQzNjkwMjksImlzcyI6IkxlYW5kcm8gTG9wZXoiLCJzdWIiOiJNYXJrZXRCaW4ifQ.KxTQ2RLkdVsy3gMV7rj0ZOUQC_0Ux9EsYRb8VuUbwrY");
  //       // },
  //     })
  //     // .done(function () {
  //         // window.location.href = "/settings";
  //     //   })
  //       .fail(function () {
  //         console.log("Failed!!!")
  //         $('.ui.error.message').empty().append("Invalid email or password").show()
  //         $(this).removeClass('loading');
  //       }).always(function () {
  //       });
  //   }
  //     $(this).removeClass('loading');
  // });

});
