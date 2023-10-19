$(document).ready(function() {
    getFormularios()
   
    $('#fromularios').DataTable();
  });
  
  
  
  
  async function getFormularios() {
    const request = await fetch("http://3.83.152.157:5000/", {
      method: 'GET',
    })
  
    let response = await request.json()
    console.log(response)
    if (request.ok) {
              let text = ``
              response.forEach(formulario => {
                  text +=
                      `
                       <tr>
                          <th>${formulario.nombre}</th>
                          <th>${formulario.informacion}</th>
                          <th>${formulario.version}</th>
                          <th>${formulario.fecha}</th>
                          <th><button onclick="redirectChecks('${formulario.id}')" class="btn btn-primary btn-user"> Editar Controles </button></th>
                      </tr>
                      `
              });
              document.getElementById("formulariosTBody").innerHTML = text;
    }
    
  }