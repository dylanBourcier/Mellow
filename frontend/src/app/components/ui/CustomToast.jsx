import { CheckCircle, AlertCircle } from 'lucide-react';
import clsx from 'clsx';

export default function CustomToast({ message, type }) {
  const isSuccess = type === 'success';

  return (
    <div
      className={clsx(
        'flex items-center gap-3 px-4 py-3 rounded-xl shadow-md w-full max-w-sm border',
        isSuccess
          ? 'bg-[hsl(269,100%,96%)] text-lavender-3 border-[hsl(269,90%,92%)]' // Lavender
          : 'bg-[hsl(348,100%,97%)] text-error border-[hsl(349,79%,91%)]' // Rose
      )}
    >
      {isSuccess ? (
        <CheckCircle size={20} className="text-lavender-3" />
      ) : (
        <AlertCircle size={20} className="text-error" />
      )}
      <p className="text-sm font-medium">{message}</p>
    </div>
  );
}
