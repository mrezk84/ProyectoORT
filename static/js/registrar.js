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

  const respuesta = await fetch("http://localhost:5000/usuarios/registrar", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
   
    body: JSON.stringify(datos),  
  })
  if (respuesta.ok){
    Swal.fire(
      'Exito!',
      'Registro correctante el usuario',
      'success'
      )
  }else{
    Swal.fire({
      title: 'Error!',
      text: 'Error al registrar usuario, verifique datos',
      icon: 'error',
    
    })

  }
  const data= await respuesta.json()

  console.log(data);

  
}
  




