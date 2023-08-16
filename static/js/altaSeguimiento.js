$(document).ready(function(obra_id) {
  cargarObra(obra_id)

  // getFormularios()
  // $('#pisos').DataTable();
});


async function cargarObra(obra_id){
  const request = await fetch("http://localhost:8080/obras/Byid", {
        method: 'GET',
        body: JSON.stringify(obra_id),
    })

    const obra = await request.json();
    console.log(obra)
    // let listadoHtml = '';
    // listadoHtml += 
    if(request.ok){
      document.querySelector('#nombre').outerHTML = `<h6 class="m-0 font-weight-bold text-primary">${obra}</h6>`;
    }
    
}


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