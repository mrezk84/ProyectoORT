$(document).ready(function() {
    // on ready

  });

  async function registrarObra() {
   let datos = {};
   
     datos.nombre = document.getElementById('txtNombre').value;

          const request = await fetch("http://localhost:8080/obras/registrar", {
              method: 'POST',
              headers: {
                  'Accept': 'application/json',
                  'Content-Type': 'application/json'
              },
              body: JSON.stringify(datos)
          })
          if (request.ok){
            Swal.fire(
              'Exito!',
              'Se guardo de la manera correcta la obra',
              'success'
              )
          }else{
            Swal.fire({
              title: 'Error!',
              text: 'Error al guardar la obra',
              icon: 'error',
            
            })
          }
  }
