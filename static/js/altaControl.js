
  async function registrarControl() {
   let datos = {};
     datos.descripcion = document.getElementById('txtDescripcion').value;
     datos.tipo = document.getElementById('txtTipo').value;

          const request = await fetch('http://34.192.187.56:5000/controles/registrar', {
              method: 'POST',
              body: JSON.stringify(datos),
              headers: {
                  'Accept': 'application/json',
                  'Content-Type': 'application/json'
              },
          })
          if (request.ok){
            Swal.fire(
              'Exito!',
              'Se guardo de la manera correcta el control',
              'success'
              )
              window.location.href = 'controles.html'
              
          }else{
            Swal.fire({
              title: 'Error!',
              text: 'Error al guardar el control',
              icon: 'error',
            
            })
          }
    
  }
