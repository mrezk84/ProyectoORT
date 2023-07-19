$(document).ready(function() {
    // on ready

  });


  async function guardarFormulario(){
    let datos = {};

    datos.nombre = document.getElementById('txtNombre').value;
    datos.informacion = document.getElementById('txtInformacion').value;
    console.log(datos);
    const request = await fetch('"http://localhost:8080/formularios/registrar"', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(datos)
    })
      if (request.status == 200){
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