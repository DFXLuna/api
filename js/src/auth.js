export async function login(c, username, password) {
    const authData = await c.pb.collection('users').authWithPassword(username, password)
        .then(
            result => result,
            err => { console.log(err) }
        );
    if (authData === undefined) {
        return
    }
    console.log("Authenticated: ", pb.authStore.isValid)
    //console.log(pb.authStore.token)
}

export function isLoggedIn(c) {
    v = c?.pb.authStore.isValid
    return (v !== false) && (v !== undefined)
}
