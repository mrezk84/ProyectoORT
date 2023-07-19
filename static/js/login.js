$(document).ready(function() {
 
 });

  
 async function iniciarSesion() {
    let datos={};
     datos.email = document.getElementById('txtEmail').value;
    datos.password = document.getElementById('txtPassword').value;
    if(datos.email==""){
      errorModal("Debe ingresar el usuario")
      return
      }
     if(datos.password==""){
      errorModal("Debe ingresar la contraseña")
      return
    }

    const response = await fetch("http://localhost:8080/usuarios/login", {
         method: 'POST',
         headers: {
          'Content-Type': 'application/json',
         },
         credentials:"include",
         body: JSON.stringify(datos),
         
       });
       
       const data = await response.json()
       Swal.fire(
        'Exito!',
        'Se inició correctamente la sesión, bienvenido',
        'success'
        )
       console.log(data);
    
       
}
           
