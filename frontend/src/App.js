import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
  const [todos, setTodos] = useState([]);
  const [newTodo, setNewTodo] = useState("");

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await axios.get('http://localhost:8080/todo');
    setTodos(res.data);
  };

  const addTodo = async () => {
    if (newTodo === "") return;
    const res = await axios.post('http://localhost:8080/todo', { title: newTodo, status: false });
    setTodos([...todos, res.data]);
    setNewTodo("");
  };

  const deleteTodo = async (id) => {
    await axios.delete(`http://localhost:8080/todo/${id}`);
    setTodos(todos.filter(todo => todo.id !== id));
  };

  return (
    <div className="App">
      <h1>ToDo List</h1>
      <input
        value={newTodo}
        onChange={(e) => setNewTodo(e.target.value)}
        type="text"
        placeholder="Add a todo"
      />
      <button onClick={addTodo}>Add</button>
      {todos.map((todo, index) => (
        <div key={index}>
          <h2>{todo.title}</h2>
          <button onClick={() => deleteTodo(todo.id)}>Delete</button>
        </div>
      ))}
    </div>
  );
}

export default App;

