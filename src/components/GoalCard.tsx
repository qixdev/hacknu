interface props {
    content: string
}
const GoalCard = ({content}:props) => {
    return (
        <div className={'rounded m-4 p-6 py-10 bg-[#55bbeb] text-white text-[20px]'} style={{boxShadow: '0px 2px 2px 1px #a1a8b0'}}>
            {content}
        </div>
    );
}

export default GoalCard;