$(document).ready(function() {
 
 });

  
 async function iniciarSesion() {
    let datos={};
     datos.email = document.getElementById('txtEmail').value;
    datos.password = document.getElementById('txtPassword').value;
    if(datos.email==""){
      Swal.fire({
        title: 'Error!',
        text: 'Debe ingresar el correo para iniciar sesión',
        icon: 'error',
      
      });
    }
     if(datos.password==""){
      Swal.fire({
        title: 'Error!',
        text: 'La contraseña no puede estár vacía',
        icon: 'error',
      
      });
    }

    const response = await fetch("http://localhost:8080/usuarios/login", {
         method: 'POST',
         headers: {
          'Content-Type': 'application/json',
         },
         credentials:"include",
         body: JSON.stringify(datos),
         
       });
    
       if (response.ok){
        Swal.fire(
          'Exito!',
          'Se guardo de la manera correcta el formulario',
          'success'
          )
      }else{
        Swal.fire({
          title: 'Error!',
          text: 'Error al guardar el formulario',
          icon: 'error',
        
        })
    
  }
}
           
