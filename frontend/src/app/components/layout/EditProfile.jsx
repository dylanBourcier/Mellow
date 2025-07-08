import React from 'react';
import Input from '../ui/Input';
import Label from '../ui/Label';
import Button from '../ui/Button';
import FileInput from '../ui/FileInput';

function EditProfile() {
    return (
        <form className='flex flex-col items-center justify-center gap-2.5'>
            <div className='flex items-center justify-center mb-12'>
            <h1 className=' text-lavender-3 text-shadow-(--text-shadow) w-full text-center'>Edit Profile</h1>
            </div>
            <div className='flex gap-2.5 w-full'>
            <Label htmlFor={"firstname"}>FirstName* :<Input id="firstname" placeholder="Enter your firstname..."></Input></Label>
            <Label htmlFor={"lastname"}>LastName* :<Input id="lastname" placeholder="Enter your lastname..."></Input></Label>
            </div>
            <div className='flex flex-col items-start gap-2.5 w-full'>
            <Label htmlFor={"username"}>Username* :</Label>
                <Input id="username" placeholder="Enter your username..."></Input>
            </div>
            <div className='flex flex-col items-start gap-2.5 w-full'>
                <Label htmlFor={"birthdate"}>Birthdate* :</Label>
                <Input
                  type="date"
                  id="birthdate"
                  name="birthdate"
                  placeholder="********"
                  required
                />
            </div>
            <div className="flex flex-col lg:flex-row w-full gap-2">
                <div className="flex-1">
                  <Label htmlFor="password" className="block mb-2">
                    Password* :
                  </Label>
                  <Input 
                  type="password" 
                  id="password" 
                  name="password"
                  placeholder="********"
                  required />
                </div>
                <div className="flex-1">
                  <Label htmlFor="confirm_password" className="block mb-2">
                    Confirm Password* :
                  </Label>
                  <Input
                    type="password"
                    id="confirm_password"
                    name="confirm_password"
                    placeholder="********"
                    required
                  />
                </div>
            </div>
                <div className='flex flex-col items-start gap-2.5 w-full'>
                    <Label htmlFor="about">About me :</Label>
                    <Input
                      type="textarea"
                      id="about"
                      name="about"
                      placeholder="Tell us about yourself..."
                      className="h-24"
                    />
                </div>
                <div className='flex flex-col items-start gap-2.5 w-full'>
                    <Label htmlFor="privacy">Privacy*:</Label>
                </div>
                <div className='flex gap-2.5 w-full'>
                    <Button isSecondary={true}>Public</Button>
                    <Button isSecondary={true}>Private</Button>
                </div>
                <div className='flex flex-col items-start gap-2.5 w-full'>
                    <Label>Choose your avatar :</Label>
                    <FileInput>Avatar</FileInput>
                </div>
                <div className='flex flex-1 align-middle justify-center gap-2.5 w-full'>
                    <Button type="submit" className="w-full">
                        Save Changes
                    </Button>
                </div>
        </form>
    );
}

export default EditProfile;