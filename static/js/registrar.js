$(document).ready(function() {
  // on ready
  
  
});


async function registrarUsuario() {
  let datos = {};

  datos.email = document.getElementById('txtEmail').value;
  datos.username = document.getElementById('txtUser').value;
  datos.password = document.getElementById('txtPassword').value;
  let repetirPassword = document.getElementById('txtRepetirPassword').value;

  if (repetirPassword != datos.password) {
    alert('La contraseña que escribiste es diferente.');
    return;
  }

    const request = await fetch("http://34.192.187.56/usuarios/registrar", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },

      body: JSON.stringify(datos),

    });
    if (request.status == 201) {
      alert("Usuario creado correctamente");
    }else{
      alert("Error registrando al usuario");
    }
    window.location.href = `index.html`;
}
  





