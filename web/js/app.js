let chart = null;

async function searchUser() {

    const username = 
        document.getElementById("username").value;

    const response =
        await fetch(`/github/${username}`);

    if(!response.ok){
        alert("Usuario no encontrado");
        return;
    }

    const data =
        await response.json();

    console.log(data);

    document.getElementById("name")
        .textContent = data.name;

    document.getElementById("login")
        .textContent = data.login;

    document.getElementById("followers")
        .textContent = "Followers: " + data.followers;

    document.getElementById("repos")
        .textContent = "Repositories total: " + data.repos;

    document.getElementById("topLanguage")
        .textContent = "Main Language: " + data.language;

    document.getElementById("topRepo")
        .textContent = "Main repo: " + data.popularRepo;
    
    console.log(JSON.stringify(data.repositoryList, null, 2));

    renderChart(data.languages);


    renderRepositories(data.repositoryList);
}

function renderChart(languages) {

    const ctx =
        document.getElementById("languageChart");

    if (chart) {
        chart.destroy();
    }

    chart = new Chart(ctx, {
        type: "pie",

        data: {
            labels: Object.keys(languages),

            datasets: [{
                data: Object.values(languages)
            }]
        }
    });
}

function renderRepositories(repositories) {

    const list =
        document.getElementById("repositoryList");

    list.innerHTML = "";

    repositories.forEach(repo => {

        const li =
            document.createElement("li");

        li.textContent =
            `${repo.name} (${repo.language}): ${repo.description}`;

        list.appendChild(li);
    });
}
