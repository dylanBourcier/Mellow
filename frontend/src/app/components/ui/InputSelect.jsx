"use client";
import React from 'react';

function InputSelect({ options, id }) {
    const handleChange = (event) => {
        console.log(event.target.value);
    };
    console.log(options);
    
    return (
        <div>
            <select
                id={id}
                name={id}
                onChange={handleChange}
                className="block w-full p-2 bg-white rounded-lg shadow-(--box-shadow) focus:outline-none focus:ring-2 focus:ring-lavender-4 focus:border-lavender-4"
            >
                {options.map((option) => (
                    <option key={option.value} value={option.value}>
                        {option.label}
                    </option>
                ))}
            </select>
            
        </div>
    );
}

export default InputSelect;