import { login } from '../auth'

function Login({ c }) {
    return (
        <div className="card w-1/3 bg-base-300 shadow-m">
            <form onSubmit={(e) => { onSubmit(e, c) }} className="card-body items-center text-center">
                <h2 className="card-title justify-self-center">Login</h2>
                <label htmlFor="uname" hidden>Username</label><input name="username" type="text" placeholder="username" className="input input-bordered w-full max-w-xs" />
                <label htmlFor="pword" hidden>Password</label><input name="password" type="password" placeholder="password" className="input input-bordered w-full max-w-xs" />
                <div className="card-actions justify-end">
                    <button type="submit" className="btn btn-primary">Login</button>
                </div>
            </form>
        </div>
    )
}

function onSubmit(e, c) {
    e.preventDefault();
    const formData = new FormData(e.target);
    login(c, formData.get("username"), formData.get("password"));
}

export default Login