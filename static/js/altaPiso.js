$(document).ready(function() {
    getObras()
});


async function getObras() {

    const request = await fetch("http://3.83.152.157:5000/obras", {
        method: 'GET',
    })

    const obras = await request.json();
    console.log(obras)
    if (request.ok) {
        let listadoHtml = '';
            for (let obra of obras) {
                listadoHtml +=
                `
                 <option value="${obra.id}">${obra.id}--${obra.nombre}</option>
                `
        };

        document.getElementById("obra").innerHTML = listadoHtml;

    }
}


async function registrarPiso() {
    let datos = {};
      numero = document.getElementById('NroPiso').value;
      datos.numero = Number(numero);
      id = document.getElementById('obra').value;
      datos.obra_id = Number(id);
           const request = await fetch("http://3.83.152.157:5000/pisos/registrar", {
               method: 'POST',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.ok) {
             alert("Piso registrado correctamente");
         }else{
             alert("Error registrando el piso");
         }
             window.location.href = `pisos.html`;
   }