import React, { useState } from 'react';
const Form = (prop: any) => {

    const [data, setData] = useState('');
    
    const handleSubmit = (e: any) => {
        e.preventDefault()
        console.log(data)
      }
    return (  
        <div className="form">
            <form onSubmit={handleSubmit}>
          <input type={data} onChange={(e) => setData(e.target.value)}/>
          <button>
            Button
          </button>
          <p>{data}</p>
        </form>
        </div>
    );
}
 
export default Form;