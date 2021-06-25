import React, {useState} from 'react';

function HelloWorld() {
    const [users, setUsers] = useState([]);

    const getPersons = () => {
        window.backend.GetUsers().then(users => {
            console.log(users);
            setUsers(users);
        });
    }
    return (
        <div className="App">
            <button onClick={() => getPersons()} type="button">
                Load from Backend
            </button>
            <table>
                <thead>
                <tr>
                    <td>Name</td>
                    <td>Email</td>
                    <td>Age</td>
                </tr>
                </thead>
                <tbody>
                {
                    users.map(user =>
                        <tr key={user.age}>
                            <td>{user.name}</td>
                            <td>{user.email}</td>
                            <td>{user.age}</td>
                        </tr>
                    )
                }
                </tbody>
            </table>
        </div>
    );
}

export default HelloWorld;
