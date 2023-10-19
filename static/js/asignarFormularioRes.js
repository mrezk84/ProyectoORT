$(document).ready(function() {
    getUsuarios()
});



async function getUsuarios() {
    const request = await fetch("http://3.83.152.157:5000/usuarios", {
        method: 'GET',
    })

    const usuarios = await request.json();
    console.log(usuarios)
    if (request.ok) {
        let listadoHtml = '';
            for (let usuario of usuarios) {
                listadoHtml +=
                `
                 <option value="${usuario.id}">${usuario.id}--${usuario.name}</option>
                `
        };

        document.getElementById("select").innerHTML = listadoHtml;

    }
}

async function asignarResponsable(){
    let datos = {};
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formularioid = searchParams.get('formulario_id');
    datos.formulario_id = Number(formularioid);


    let usuario = document.getElementById('select').value;
    datos.usuario_id = Number(usuario);

    console.log(formularioid)
    console.log(usuario)

    const request = await fetch("http://34.192.187.56:5000/formularios/addUser", {
               method: 'POST',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.status == 201) {
             alert("control agregado correctamente");
         }else{
             alert("Error agregando el control");
         }
             window.location.href = `formularios.html`;
}