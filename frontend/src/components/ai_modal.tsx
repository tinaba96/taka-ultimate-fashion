import React, { useEffect, useState } from 'react';
import Modal from 'react-modal';

interface Product {
    ID: number;
    Name: string;
    ImageURL: string;
    Price: string;
}

interface SaveResponse {
  statusCode: number;
  body: {
    products: Product[];
  };
}

interface AIModalProps {
    showModal: boolean;
    setShowModal: React.Dispatch<React.SetStateAction<boolean>>;
  }

// const apiUrl = process.env.REACT_APP_API_URL;

const AIModal: React.FC<AIModalProps> = ({ showModal, setShowModal }) => {
//   const [showModal, setShowModal] = useState<boolean>(false);
const [isLoading, setIsLoading] = useState<boolean>(true);
  const [ai, setAi]  = useState<Product[]>([])

  useEffect(() => {
    // fetch(`${apiUrl}/getSave`)
    fetch('https://a4qu0cj142.execute-api.us-east-1.amazonaws.com/tuf-ai')
    .then(response => response.json())
    .then((data: SaveResponse) => {
      console.log('data.body')
      console.log(data.body)
      setIsLoading(false);
        // // var saveProducts = data.datas.map((d: Save) => d.Product)
      setAi(data.body.products)
    })
    .catch(error => {
      setIsLoading(false);
        console.error('Request Error(Please check the response type):', error);
    });
},[])


  const handleDeleteAction = () => {
    setShowModal(false);
  };


  const closeModal = () => setShowModal(false);

  const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      transform: 'translate(-50%, -50%)',
      maxHeight: '80vh',
      padding: '20px',
    },
  };

  const confirmModal = (
    <Modal
      style={customStyles}
          isOpen={showModal}
      onRequestClose={closeModal}>
        <h1 className="pt-10 text-2xl header-title">AI Recommend</h1>
        <p className="pb-1">This is an AI recommendation based on the 'Favorites' items.</p>
        {isLoading && <p className='pb-10'>Please wait for AI...</p>}
        
            <div  className="grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
            {ai.length !== 0 ? (
              ai.map((product) => (
                  <div key={product.ID} className="group relative">
                  <div className="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
                      <img
                        alt={product.Name}
                        src={`${product.ImageURL}`}
                        className="h-full w-full object-cover object-center lg:h-full lg:w-full"
                      />
                  </div>
                  <div className="mt-4 flex justify-between">
                      <div>
                      <h3 className="text-sm text-gray-700">
                          <a href={product.ImageURL}>
                          {product.Name}
                          </a>
                      </h3>
                      <p className="text-sm font-medium text-gray-900">{product.Price}</p>
                      </div>
                  </div>
                  </div>
              ))
            ) : (
              <div className='text-gray-900 font-bold header-title'>No products available</div>
            )}
            </div>
            <div className="px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
        <button
          type="button"
          className="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
          onClick={handleDeleteAction}>
          CLOSE
        </button>
      </div>
      
    </Modal>
  );

  return (
    <>
      {confirmModal}
    </>
  );
}
export default AIModal;
