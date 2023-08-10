$(document).ready(function() {
    // on ready
    
  });
  
  
  async function enviarMail() {
   let datos = {};
   datos.mail = document.getElementById('txtMail').value;

    
   const respuesta = await fetch('"http://localhost:8080/enviarCorreo', 
   {
     method: 'POST',
     headers: {
       'Accept': 'application/json',
       'Content-Type': 'application/json'
     },
     body: JSON.stringify(datos)
  });
  if (respuesta.ok){

    Swal.fire(
      'Exito!',
      'Se envió correo a su casilla para cambiar la contraseña',
      'success'
      )
    window.location.href = 'index.html'
  } else{
    Swal.fire({
      title: 'Error!',
      text: 'Error al enviar el correo',
      icon: 'error',
    
    })
  }

} 
  
