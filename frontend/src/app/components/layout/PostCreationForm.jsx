import React from 'react';
import InputSelect from '../ui/InputSelect';
import Label from '../ui/Label';
import Input from '../ui/Input';

function PostCreationForm() {
    return (
        <form className='flex flex-col gap-2.5 w-full max-w-[600px]'>
            <div>
                <Label htmlFor={"visibility-select"}>Post On* :</Label>
                <InputSelect></InputSelect>
            </div>
            <div>
                <Label htmlFor={"title"}>Title* :</Label>
                <Input placeholder="Enter your title here ..." id="title"></Input>
            </div>
            <div>
                <Label htmlFor={"content"}>Content* :</Label>
                <Input></Input>
            </div>
            <div>
                <Label htmlFor={"image"}></Label>

            </div>
            <div>
                <Label htmlFor={"visibility"}>Visibility* :</Label>

            </div>
            <div>
                <Label htmlFor={""}></Label>
                <InputSelect></InputSelect>
            </div>
        </form>
    );
}

export default PostCreationForm;