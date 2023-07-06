$(document).ready(function() {
    cargarObras();
  $('#controles').DataTable();


});

async function cargarObras() {
    const request = await fetch('/obra', {
      method: 'GET',
      headers: getHeaders()
    });
    const obras = await request.json();

    let listadoHtml = '';
    for (let obra of obras) {
        let formularioHtml = '<option value="'+ obra +'">'+ obra.nombre +'</option>';
        listadoHtml += formularioHtml;
      }
  

      document.querySelector('#obra').outerHTML = listadoHtml;
  }


  async function buscar() {
    let obra = document.getElementById('obra').value;
    const request = await fetch('/check/buscar' + obra, {
      method: 'GET',
      headers: getHeaders()
    });
    const controles = await request.json();

    let listadoHtml = '';
    for (let control of controles) {
      let controlHtml = '<tr><td>'+ control.id+'</td><td>' + control.descripcion + '</td><td>' + control.fecha + '</td><td>'
                  + control.formulario.id +'</td><td>';
      listadoHtml += controlHtml;
  }


document.querySelector('#controles tbody').outerHTML = listadoHtml;
  }


function getHeaders() {
    return {
     'Accept': 'application/json',
     'Content-Type': 'application/json',
       'Authorization': localStorage.token  
   };
  }
  