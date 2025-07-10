"use client"

import React from 'react';
import InputSelect from '../ui/InputSelect';
import Label from '../ui/Label';
import Input from '../ui/Input';
import FileInput from '../ui/FileInput';
import Button from '../ui/Button';
import { useForm } from 'react-hook-form';
import SelectFollowers from '../ui/SelectFollowers';
import { Controller } from 'react-hook-form';

function PostCreationForm() {
    const { control, handleSubmit, formState: { errors } } = useForm({
        defaultValues: {
            selectedFollowers: [] // ou la valeur par défaut appropriée
        }
    });
    const {setValue}= useForm();
    
    const whoCanSeeOptions=[
        { value: 'everyone', label: 'Everyone' },
        { value: 'group', label: 'group' },
        { value: 'group2', label: 'group2' },
    ]

    const followers = [
        { id: 1, name: 'John Doe' },
        { id: 2, name: 'Jane Smith' },
        { id: 3, name: 'Alice Johnson' },
        { id: 4, name: 'Bob Brown' },
    ]
    return (
        <form className='flex flex-col gap-6 w-full max-w-[600px]'>
            <div className='flex flex-col gap-2.5'>
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
                <Input placeholder="Enter your content here ... " type='textarea'></Input>
            </div>
            <div>
                <Label htmlFor={"image"}></Label>

            </div>
            <div>
                <Label htmlFor={"image"}>Image :</Label>
                <FileInput id="image" accept="image/*" usePreview={false} setValue={setValue}className='w-full' />
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
                <div className='flex flex-col gap-2.5'>
                    <Label htmlFor={""}>Who can see your post* :</Label>
                    <Controller
                        name="selectedFollowers"
                        control={control}
                        render={({ field }) => (
                            <SelectFollowers 
                                followers={followers}
                                value={field.value}
                                onChange={field.onChange}
                            />
                        )}
                    />
                </div>
                </div>
                <Button type="submit" className='w-full' isSecondary={false}>
                Post
                </Button>
        </form>
    );
}

export default PostCreationForm;