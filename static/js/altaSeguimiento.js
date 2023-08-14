async function registrarSeguimiento() {
    let datos = {};
      datos.pendiente = document.getElementById('txt').value;
      datos.pendiente2 = document.getElementById('txt').value;
 
           const request = await fetch('http://localhost:8080/', {
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
               'Se guardo de la manera correcta el seguimiento',
               'success'
               )
               window.location.href = ''
               
           }else{
             Swal.fire({
               title: 'Error!',
               text: 'Error al guardar el seguimiento',
               icon: 'error',
             
             })
           }
     
   }