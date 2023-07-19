$(document).ready(function() {
  // on ready
  
  
});


async function registrarUsuario() {
  let datos = {};

  datos.email = document.getElementById('txtEmail').value;
  datos.nombre = document.getElementById('txtUser').value;
  datos.password = document.getElementById('txtPassword').value;
  let repetirPassword = document.getElementById('txtRepetirPassword').value;

  if (repetirPassword != datos.password) {
    alert('La contraseña que escribiste es diferente.');
    return;
  }

  const respuesta = await fetch("http://localhost:8080/usuarios/registrar", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },

    body: JSON.stringify(datos),
    
  });

  const data = await respuesta.text()

  console.log(data);
 
  }
  





