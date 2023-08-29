$(document).ready(function() {
  cargarObra()

  $('#obra').DataTable();
  // getFormularios()
  // $('#pisos').DataTable();
});

let id = null;

async function cargarObra(){
  const url = new URL(document.URL);
  const searchParams = url.searchParams;

  id = searchParams.get('obra_id');

  const request = await fetch("http://localhost:8080/obras/Byid", {
        method: 'GET',
        body: JSON.stringify(id),
    })

    const obra = await request.json();
    console.log(obra)
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