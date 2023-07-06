$(document).ready(function() {
    // on ready

  });


  async function guardarFormulario(){
    let datos = {};

    datos.nombre = document.getElementById('txtNombre').value;
    datos.informacion = document.getElementById('txtInformacion').value;
    console.log(datos);
    const request = await fetch('api/v1/formulario/guardar', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(datos)
    })
      if (request.status == 200){
          alert("Formulario guardado correctamente")
      }else{
            alert("Error al guardar el formulario")
      }


  }