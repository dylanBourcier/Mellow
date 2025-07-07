import React from 'react';
import InputSelect from '../ui/InputSelect';
import Label from '../ui/Label';
import Input from '../ui/Input';

function PostCreationForm() {
    const visibilityOptions=[
        { value: 'public', label: 'Public' },
        { value: 'private', label: 'Private' },
        { value: 'almost-private', label: 'Almost Private' },
    ]
    const whoCanSeeOptions=[
        { value: 'everyone', label: 'Everyone' },
        { value: 'group', label: 'group' },
        { value: 'group2', label: 'group2' },
    ]
    return (
        <form className='flex flex-col gap-2.5 w-full max-w-[600px]'>
            <div>
                <Label htmlFor={"postOn"}>Post On* :</Label>
                <InputSelect id="postOn" options={whoCanSeeOptions}></InputSelect>
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
                <Label htmlFor={"visibility-select"}></Label>
                <div className='flex gap-1 justify-center'>
                    <div className='flex-1'>
                    <input type='radio' name="visibility" id="public" className='hidden peer'/>
                    <label htmlFor="public" className='flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200'>Public</label>
                    </div>
                    <div className='flex-1'>
                    <input type='radio' name="visibility" id="almost-private" className='hidden peer'/>
                    <label htmlFor="almost-private" className='flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200'>Almost Private</label>
                    </div>
                    <div className='flex-1'>
                    <input type='radio' name="visibility" id="private" className='hidden peer'/>
                    <label htmlFor="private" className='flex align-middle justify-center py-2 border border-lavender-5 text-lavender-5 bg-transparent rounded-2xl peer-checked:bg-lavender-6 cursor-pointer transition-colors duration-200'>Private</label>
                    </div>
                </div>
            </div>
            <div>
                <Label htmlFor={"image"}>Image :</Label>
                <Input type="file" id="image" accept="image/*"></Input>
            </div>
            {/* <div>
                <Label htmlFor={""}>Who can see your post* :</Label>
                <InputSelect></InputSelect>
            </div> */}
        </form>
    );
}

export default PostCreationForm;