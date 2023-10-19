$(document).ready(function() {

   
    Paginar("auditoria");
  });


  async function getFormularios() {
    const request = await fetch("http://34.192.187.56:8080/formulario", {
      method: 'GET',
    })
  
    let response = await request.json()
    console.log(response)
    if (request.ok) {
              let text = ``
              response.forEach(auditoria => {
                  text +=
                      `
                       <tr>
                          <th>${auditoria.nombre}</th>
                          <th>${auditoria.informacion}</th>
                          <th>${auditoria.version}</th>
                          <th>${auditoria.fecha}</th>
                          
                      </tr>
                      `
              });
              document.getElementById("auditoriaTBody").innerHTML = text;
    }
    
  }