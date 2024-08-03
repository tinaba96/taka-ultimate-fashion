import React, {ChangeEvent, useEffect, useState} from 'react';
import Select from "react-tailwindcss-select";
import AIModal from '../components/ai_modal';
import { runInContext } from 'vm';

const options: Option[] = [
    { value: "JEANS", label: "JEANS" },
    { value: "T-SHIRTS", label: "T-SHIRTS" },
    { value: "SWIMWEAR", label: "SWIMWEAR" },
    { value: "SHIRTS", label: "SHIRTS" },
    { value: "SHORTS", label: "SHORTS" },
    { value: "PANTS", label: "PANTS" },
    { value: "GRAPHIC T-SHIRTS", label: "GRAPHIC T-SHIRTS" },
    { value: "ACTIVEWEAR", label: "ACTIVEWEAR" },
];
interface Option {
    value: string;
    label: string;
}

  interface Product {
    ID: number;
    Name: string;
    // href: string;
    ImageURL: string;
    // imageAlt: string;
    Price: string;
    // color: string;
  }

  interface DataResponse {
    datas: Product[];
}
interface Save {
  ID: number;
  Status: number
  ProductID: number;
  Product: Product;
}

interface HeartResponse {
  datas: Save[];
}
type Dictionary = {
  [key: number]: number;
};

const apiUrl = process.env.REACT_APP_API_URL;

const Home: React.FC = () => {
  const [showModal, setShowModal] = useState<boolean>(false);
    const [products, setProducts] = useState<Product[]>([])
  const [heart, setHeart] = useState<Dictionary>({})
    // const [dropdownOpen, setDropdownOpen] = useState(false);
    const [categories, setCategories] = useState<Option[]>([]);
    const [minPrice, setMinPrice] = useState<number>(0);
    const [maxPrice, setMaxPrice] = useState<number>(0);
    // const [categories, setCategories] = useState<string[]>([]);

    useEffect(() => {
      fetch(`${apiUrl}/getSave`)
      .then(response => response.json())
      .then((data: HeartResponse) => {
          var saveProducts = data.datas
          const dictionary = saveProducts.reduce<Dictionary>((acc, item) => {
            acc[item.ProductID] = item.Status;
            return acc;
          }, {});
          setHeart(dictionary)
      })
      .catch(error => {
          console.error('Request Error(Please check the response type):', error);
      });
  },[])


  const handleMinChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = parseFloat(e.target.value);
    if (value >= 0 && value <= maxPrice) {
      setMinPrice(value);
    }
  };

  const handleMaxChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = parseFloat(e.target.value);
    if (value >= minPrice) {
      setMaxPrice(value);
    }
  };
  const searchByPrice = () => {
    const selectedCategories = categories.map(option => option.value)
      fetch(`${apiUrl}/products?categories=${selectedCategories}&minPrice=${minPrice}&maxPrice=${maxPrice}`)
        .then(response => response.json())
        .then((data: DataResponse) => {
            // console.log(data)
            setProducts(data.datas != null ? data.datas : [])
        })
        .catch(error => {
            console.error('Request Error(Please check the response type):', error);
        });
  }
    const handleChange = (selectedOption: Option | Option[] | null) => {
      setCategories(selectedOption as Option[]);
      const selectedCategories = (Array.isArray(selectedOption) && selectedOption !== null && selectedOption.length !== 0) ? selectedOption.map(option => option.value) : []
      // setCategories(selectedCategories as string[])
      fetch(`${apiUrl}/products?categories=${selectedCategories}&minPrice=${minPrice}&maxPrice=${maxPrice}`)
        .then(response => response.json())
        .then((data: DataResponse) => {
            setProducts(data.datas != null ? data.datas : [])
            // this.setState({ products: response.data });
        })
        .catch(error => {
            console.error('Request Error(Please check the response type):', error);
        });
      }

      const handleClick = async (e: React.MouseEvent<HTMLImageElement>, product_id:number) => {
        e.preventDefault();
        
        const data = {
          ProductID: product_id,
          Status: heart[product_id] === 1 ? 0 : 1
        };
    
        try {
          const res = await fetch(`${apiUrl}/save`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
          });
    
          await res.json();
        } catch (error) {
          console.error('Error:', error);
        }


        // update heat
        fetch(`${apiUrl}/getSave`)
        .then(response => response.json())
        .then((data: HeartResponse) => {
            var saveProducts = data.datas
            const dictionary = saveProducts.reduce<Dictionary>((acc, item) => {
              acc[item.ProductID] = item.Status;
              return acc;
            }, {});
            setHeart(dictionary)
        })
        .catch(error => {
            console.error('Request Error(Please check the response type):', error);
        });
      };

    // const toggleDropdown = () => {
    //     setDropdownOpen(!dropdownOpen);
    // };
    // const headerText = categories.map(category => category.value).join(' & ');
  if (!categories) {
    setCategories([])
    }
    const headerText = (categories !== null && categories?.length !== 0) ? categories.map(category => category.value).join(' & ') : "ALL";

    useEffect(() => {
      fetch(`${apiUrl}/products`)
        .then(response => response.json())
        .then((data: DataResponse) => {
            setProducts(data.datas)
            // console.log(data)
        })
        .catch(error => {
            console.error('Request Error(Please check the response type):', error);
        });
    }, [])

  return (
    <div>
      
      
    



      <AIModal showModal={showModal} setShowModal={setShowModal} />
        <div className="max-w-2xl mx-auto py-0 sm:px-6 sm:py-16 lg:max-w-7xl lg:px-8">
            <div className="flex items-center space-x-4">
                <span>SEARCH</span>
                <Select
                    primaryColor={"indigo"}
                    value={categories}
                    onChange={handleChange}
                    options={options}
                    isMultiple={true}
                />
                <button onClick={() => setShowModal(true)} className="bg-green-700 hover:bg-green-800 text-white font-bold py-2 px-4 rounded">
                      🧠 AI 🤖 Recommend
                </button>
            </div>

      <div className="w-full flex justify-between my-6">
        <div className="flex w-1/3 h-[45px] items-center gap-2 flex-grow px-8">
          <span className="">Min</span>
          <input type="number" className="w-full text-center border rounded-md border-primary-mid" value={minPrice} id="minimum" step="0.01" min="0" onChange={(e) => handleMinChange(e)}/>
        </div>
        <div className="flex items-center text-center justify-center px-3"><p>-</p></div>
        <div className="flex w-1/3 items-center gap-2 flex-grow px-8">
          <span>Max</span>
          <input type="number" className="text-center w-full border rounded-md border-primary-mid" value={maxPrice}  step="0.01" min="0" onChange={(e) => handleMaxChange(e)}/>
        </div>
        <button onClick={searchByPrice} className="space-x-4 bg-gray-700 hover:bg-green-800 text-white font-bold py-2 px-4 rounded">
                      Price Search
        </button>
      </div>



        <h1 className="py-10 text-2xl header-title">{products.length !== 0 ? headerText : ''}</h1>
            <div  className="grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
            {products.length !== 0 ? (
              products.map((product) => (
                  <div key={product.ID} className="group relative">
                  <div className="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
                      <img
                      alt={product.Name}
                      src={`${product.ImageURL}`}
                      className="h-full w-full object-cover object-center lg:h-full lg:w-full"
                      onClick={(e) => handleClick(e, product.ID)}
                      />
                  </div>
                  <div className="mt-4 flex justify-between">
                      <div>
                      <h3 className="text-sm text-gray-700">
                        <a href={`${product.ImageURL}`}>
                          {product.Name}
                          </a>
  <svg className="h-5 w-5 text-red-500"  viewBox="0 0 24 24"  fill={heart[product.ID] === 1 ? "currentColor" : "none"}  stroke="currentColor"  strokeWidth="2"  strokeLinecap="round"  strokeLinejoin="round">  <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z" /></svg>
                      </h3>
                      <p className="text-sm font-medium text-gray-900">CA${product.Price}</p>
                      </div>
                  </div>
                  </div>
              ))
            ) : (
              <div className='text-gray-900 font-bold header-title'>No products available</div>
            )}
            </div>
        </div>
          </div>
  );
};

export default Home;
