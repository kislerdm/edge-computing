function start(r, g, b) {
    let n = N(r, g, b);
    n = n === "" ? "Not found" : n;
    const t = T(r, g, b) === true ? "Warm" : "Cool";
    document.getElementById("color_output").innerText = `<div><label for="output_name" id="output_label">Color Name:</label><output name="color_name" id="output_name"> ` +
        n + `</output></div><div><label for="output_type" id="output_label">Color Type:</label><output name="color_type" id="output_type"> ` + t + `</output></div>`;
}
