"use client";
import React from 'react';

function InputSelect() {
    const options = [
        {value: 'Everyone', label: 'Everyone'
        },
        {value: 'Group', label: 'Group'},
        {value: 'Group Admin', label: 'Group Admin'},
        {value: 'Group Members', label: 'Group Members'},
    ]
    const handleChange = (event) => {
        console.log(event.target.value);
    };
    return (
        <div>
            <select
                id="visibility-select"
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