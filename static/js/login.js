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

    const response = await fetch("http://localhost:5000/usuarios/login", {
         method: 'POST',
         headers: {
          'Content-Type': 'application/json',
         },
         body: JSON.stringify(datos),
         
       });
    
       if (response.ok){
        Swal.fire(
          'Exito!',
          'Se inicio correctamente la sesión',
          'success'
          )
          window.location.href = 'obras.html'
      }else{
        Swal.fire({
          title: 'Error!',
          text: 'Credenciales invalidas',
          icon: 'error',
        
        })
    
  }
}
           
