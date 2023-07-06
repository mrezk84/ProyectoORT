$(document).ready(function() {
    // on ready
    
  });
  
  
  async function enviarMail() {
   let datos = {};
   datos.cedula = document.getElementById('txtCedula').value;
   datos.mail = document.getElementById('txtMail').value;

    
   const respuesta = await fetch('api/enviarCorreo', 
   {
     method: 'POST',
     headers: {
       'Accept': 'application/json',
       'Content-Type': 'application/json'
     },
     body: JSON.stringify(datos)
  });
  if (respuesta.status=200){

   
    localStorage.datos= datos.cedula
    localStorage.datos= datos.mail
    alert("Se le ha enviado un código para cambiar la contraseña")
    window.location.href = 'index.html'
  }

} 
  
