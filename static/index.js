var editor = ace.edit("program-code");
editor.session.setMode("ace/mode/assembly_x86");
editor.setTheme("ace/theme/one_dark");
editor.setOptions({
    fontSize: "14pt"
});
editor.renderer.setPadding(20);
editor.setValue(
    "; ==== TUTORIAL ====\n;\n; 1. Load an assembly program (.asm)\n; 2. Make the changes in the editor\n; 3. Click the 'Load Buffer' button to load changes\n; 4. Click the 'run' button to run the program\n;\n; === END OF TUTORIAL ===",
    1
);

$(document).ready(function(){
    $('.ide.menu .item').tab({history:false});
    setRegisters(new Array(32).fill(0));
    setMemoryBlocks(new Array(1000).fill(0));
});

function setRegisters(regs) {
    var tbody = $("#register-container tbody");
    tbody.html("");
    for (var i = 0; i < 4; i++) {
        var tr = document.createElement("tr");
        for (var j = 0; j < 8; j++) {
            const idx = i * 8 + j;
            let val = document.createElement("span");
            val.style = "float: right";
            val.appendChild(document.createTextNode(regs[idx]));
            var td = document.createElement("td");
            td.appendChild(document.createTextNode(`r${idx}: `));
            td.appendChild(val);
            tr.appendChild(td);
        }
        tbody.append(tr);
    }
}

function setMemoryBlocks(mem) {
    var tbody = $("#memory-container tbody");
    tbody.html("");
    for (var i = 0; i < 100; i++) {
        var tr = document.createElement("tr");
        for (var j = 0; j < 10; j++) {
            const idx = i * 8 + j;
            const val = mem[idx];
            var td = document.createElement("td");
            td.appendChild(document.createTextNode(val));
            tr.appendChild(td);
        }
        tbody.append(tr);
    }
}

function loadFileContents(file) {
    if (file) {
        var reader = new FileReader();
        reader.readAsText(file, "UTF-8");
        reader.onload = function (evt) {
            editor.setValue(evt.target.result, -1);

            toastSuccess("Success!", "The file's content has been loaded sucessfully");
            return evt.target.result;
        };
        reader.onerror = function (evt) {
            document.getElementById("program-code").innerHTML = "Error reading file";

            $("body").toast({
                title: "Error!",
                class: "error",
                message: "There was a problem while loading the file...",
                showProgress: "bottom",
                classProgress: "black",
                position: "bottom right",
                displayTime: 5000
            });
            return "Error reading file";
        };
    }
}

function consoleHelloWorld() {
    printConsole("[+] Assembly program ready !");
}

function scrollBottomConsole() {
    var console = $("#console");
    if (console.length) psconsole.scrollTop(psconsole[0].scrollHeight - psconsole.height());
}

function printConsole(message) {
    $("#console").append(message + "\n");
    scrollBottomConsole();
}

function unlockButtons() {
    $("#btn-step").removeClass("disabled");
    $("#btn-run").removeClass("disabled");
}

async function updateStats(nbLines, nbCycles, pc) {
    $("#nb-lines span.stat").html(nbLines);
    $("#nb-cycles span.stat").html(nbCycles);
    $("#pc span.stat").html(pc);
}

$("#mainprogramfile").change(() => {
    var file = document.getElementById("mainprogramfile").files[0];
    $("#mainprogramfile-text").html(file.name);

    loadFileContents(file);
});

$("#memoryfile").change(() => {
    var file = document.getElementById("mainprogramfile").files[0];
    $("#memoryfile-text").html(file.name);
});

$("#btn-loadbuffer").click(() => {
    sendProgramContent(editor.getValue());
});

$("#btn-run").click(() => {
    runCode();
});