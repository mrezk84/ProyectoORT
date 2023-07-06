$(document).ready(function() {
    buildControles()
});



let formulario = null;
async function buildControles() {
    const url = new URL(document.URL);
    const searchParams = url.searchParams;

    formulario = searchParams.get('formulario');
    formulario = JSON.parse(decodeURIComponent(formulario));

    let text = ``;
    formulario.control.forEach(control => {
        text +=
            `
            <tr>
                <th>${control.descripcion}</th>
            </tr>
            `
    });
    document.getElementById("controlesTBody").innerHTML = text;
}

function redirectAgregarControl(){
    window.location.href = `agregarControlesFormulario.html?formulario=${encodeURIComponent(JSON.stringify(formulario))}`;
}