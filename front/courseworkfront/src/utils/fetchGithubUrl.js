



const fetchGithubUrl = (from) => {
    const rootUrl = 'https://github.com/oauth/authorize';
    const options = {
        client_id:'b76f24e4391dc2f6fdf7',
        redirect_url: 'http://localhost:80/oauth/sign-in',
        scope: 'user:email',
        state: from
    }

    const qs = new URLSearchParams(options)

    return `${rootUrl}?${qs.toString()}`
}


export default fetchGithubUrl