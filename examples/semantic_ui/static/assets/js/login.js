$(document).ready(function () {

  $('.ui.form').form({
      fields: {
        email: 'empty',
        password: 'empty'
      }
    });

  // $('.input').on('click', function () {
  //   $('.ui.error.message').empty()
  // });

  $("#submit").on('click', function () {

      var _email, _password
      $(this).addClass('loading');
      var $form = $('.ui.form')
      _email = $form.form('get value', 'email')
      _password = $form.form('get value', 'password')

    if (_email != '' && _password != '') {
    //   $('.ui.error.message').empty().append("Email and Passwor must have value").show()
    // } else {

      var data = {
        email: _email,
        password: _password
      };

      var jqxhr = $.ajax({
        method: "POST",
        url: "/login",
        data: data,
      })
        .fail(function () {
          // console.log("Failed!!!")
          // $('.ui.error.message').empty()
          // $('.ui.error.message').append('<ul class="list"><li>Invalid email or password</li></ul>')
          $(this).removeClass('loading');
        }).always(function () {
        });
    }
      $(this).removeClass('loading');
  });

});
