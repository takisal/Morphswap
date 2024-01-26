// SPDX-License-Identifier: MIT

// MorphSwap

pragma solidity ^0.8.0;


import "./IERC20.sol";
import "./extensions/IERC20Metadata.sol";
import "./utils/Context.sol";

/**
 * @dev Implementation of the {IERC20} interface.
 *
 * This implementation is agnostic to the way tokens are created. This means
 * that a supply mechanism has to be added in a derived contract using {_mint}.
 * For a generic mechanism see {ERC20PresetMinterPauser}.
 *
 * TIP: For a detailed writeup see our guide
 * https://forum.zeppelin.solutions/t/how-to-implement-erc20-supply-mechanisms/226[How
 * to implement supply mechanisms].
 *
 * We have followed general OpenZeppelin Contracts guidelines: functions revert
 * instead returning `false` on failure. This behavior is nonetheless
 * conventional and does not conflict with the expectations of ERC20
 * applications.
 *
 * Additionally, an {Approval} event is emitted on calls to {transferFrom}.
 * This allows applications to reconstruct the allowance for all accounts just
 * by listening to said events. Other implementations of the EIP may not emit
 * these events, as it isn't required by the specification.
 *
 * Finally, the non-standard {decreaseAllowance} and {increaseAllowance}
 * functions have been added to mitigate the well-known issues around setting
 * allowances. See {IERC20-approve}.
 */

contract AssetPool is Context, IERC20, IERC20Metadata {
    event Redeem(address indexed from, uint256 value, uint256 lpsent);
    mapping(address => uint256) private _balances;

    mapping(address => mapping(address => uint256)) private _allowances;

    uint256 private _totalSupply;

    string private _name;
    string private _symbol;
    address public _poolAsset;
    address public _overallContract;
    uint public _pairID;
    bool public _isNativeCoin;
    uint oneQuadrillion;
    bool public _isTipPool;
    struct queueEntry {
        uint currentBlockNum;
        uint sentLP;
    }
    mapping(address => queueEntry[]) public _queueEntryMapping;
    IERC20 _chain1AssetInterface;

    /**
     * @dev Sets the values for {name} and {symbol}.
     *
     * The default value of {decimals} is 18. To select a different value for
     * {decimals} you should overload it.
     *
     * All two of these values are immutable: they can only be set once during
     * construction.
     */
    constructor(address c1a, uint pid, bool istippool) {
        _poolAsset = c1a;
        _name = "MorphSwap LP";
        _symbol = "MSLP";
        _totalSupply = 0;
        _overallContract = msg.sender;
        _chain1AssetInterface = IERC20(c1a);
        _pairID = pid;
        oneQuadrillion = 1000000000000000;
        if (c1a == address(0)) {
            _isNativeCoin = true;
        } else {
            _isNativeCoin = false;
        }
        _isTipPool = istippool;
        //_approve(address(this), address(this), type(uint256).max);
    }

    /**
     * @dev Returns the name of the token.
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the symbol of the token, usually a shorter version of the
     * name.
     */
    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns the number of decimals used to get its user representation.
     * For example, if `decimals` equals `2`, a balance of `505` tokens should
     * be displayed to a user as `5.05` (`505 / 10 ** 2`).
     *
     * Tokens usually opt for a value of 18, imitating the relationship between
     * Ether and Wei. This is the value {ERC20} uses, unless this function is
     * overridden;
     *
     * NOTE: This information is only used for _display_ purposes: it in
     * no way affects any of the arithmetic of the contract, including
     * {IERC20-balanceOf} and {IERC20-transfer}.
     */
    function decimals() public view virtual override returns (uint8) {
        return 18;
    }

    /**
     * @dev See {IERC20-totalSupply}.
     */
    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev See {IERC20-balanceOf}.
     */
    function balanceOf(
        address account
    ) public view virtual override returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev See {IERC20-transfer}.
     *
     * Requirements:
     *
     * - `recipient` cannot be the zero address.
     * - the caller must have a balance of at least `amount`.
     */
    function transfer(
        address recipient,
        uint256 amount
    ) public virtual override returns (bool) {
        _transfer(_msgSender(), recipient, amount);
        return true;
    }

    /**
     * @dev See {IERC20-allowance}.
     */
    function allowance(
        address owner,
        address spender
    ) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev See {IERC20-approve}.
     *
     * NOTE: If `amount` is the maximum `uint256`, the allowance is not updated on
     * `transferFrom`. This is semantically equivalent to an infinite approval.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function approve(
        address spender,
        uint256 amount
    ) public virtual override returns (bool) {
        _approve(_msgSender(), spender, amount);
        return true;
    }

    /**
     * @dev See {IERC20-transferFrom}.
     *
     * Emits an {Approval} event indicating the updated allowance. This is not
     * required by the EIP. See the note at the beginning of {ERC20}.
     *
     * NOTE: Does not update the allowance if the current allowance
     * is the maximum `uint256`.
     *
     * Requirements:
     *
     * - `sender` and `recipient` cannot be the zero address.
     * - `sender` must have a balance of at least `amount`.
     * - the caller must have allowance for ``sender``'s tokens of at least
     * `amount`.
     */
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public virtual override returns (bool) {
        uint256 currentAllowance = _allowances[sender][_msgSender()];
        if (currentAllowance != type(uint256).max) {
            require(
                currentAllowance >= amount,
                "ERC20: transfer amount exceeds allowance"
            );
            unchecked {
                _approve(sender, _msgSender(), currentAllowance - amount);
            }
        }

        _transfer(sender, recipient, amount);

        return true;
    }

    /**
     * @dev Atomically increases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function increaseAllowance(
        address spender,
        uint256 addedValue
    ) public virtual returns (bool) {
        _approve(
            _msgSender(),
            spender,
            _allowances[_msgSender()][spender] + addedValue
        );
        return true;
    }

    /**
     * @dev Atomically decreases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     * - `spender` must have allowance for the caller of at least
     * `subtractedValue`.
     */
    function decreaseAllowance(
        address spender,
        uint256 subtractedValue
    ) public virtual returns (bool) {
        uint256 currentAllowance = _allowances[_msgSender()][spender];
        require(
            currentAllowance >= subtractedValue,
            "ERC20: decreased allowance below zero"
        );
        unchecked {
            _approve(_msgSender(), spender, currentAllowance - subtractedValue);
        }

        return true;
    }

    /**
     * @dev Moves `amount` of tokens from `sender` to `recipient`.
     *
     * This internal function is equivalent to {transfer}, and can be used to
     * e.g. implement automatic token fees, slashing mechanisms, etc.
     *
     * Emits a {Transfer} event.
     *
     * Requirements:
     *
     * - `sender` cannot be the zero address.
     * - `recipient` cannot be the zero address.
     * - `sender` must have a balance of at least `amount`.
     */
    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal virtual {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");

        _beforeTokenTransfer(sender, recipient, amount);

        uint256 senderBalance = _balances[sender];
        require(
            senderBalance >= amount,
            "ERC20: transfer amount exceeds b3alance"
        );
        unchecked {
            _balances[sender] = senderBalance - amount;
        }
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);

        _afterTokenTransfer(sender, recipient, amount);
    }

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     */

    /**
     * @dev Destroys `amount` tokens from `account`, reducing the
     * total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function _burn(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: burn from the zero address");

        _beforeTokenTransfer(account, address(0), amount);

        uint256 accountBalance = _balances[account];
        require(
            accountBalance >= amount,
            "ERC20: burn amount exceeds b4alance"
        );
        unchecked {
            _balances[account] = accountBalance - amount;
        }
        _totalSupply -= amount;

        emit Transfer(account, address(0), amount);

        _afterTokenTransfer(account, address(0), amount);
    }

    /**
     * @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
     *
     * This internal function is equivalent to `approve`, and can be used to
     * e.g. set automatic allowances for certain subsystems, etc.
     *
     * Emits an {Approval} event.
     *
     * Requirements:
     *
     * - `owner` cannot be the zero address.
     * - `spender` cannot be the zero address.
     */
    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) internal virtual {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Hook that is called before any transfer of tokens. This includes
     * minting and burning.
     *
     * Calling conditions:
     *
     * - when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
     * will be transferred to `to`.
     * - when `from` is zero, `amount` tokens will be minted for `to`.
     * - when `to` is zero, `amount` of ``from``'s tokens will be burned.
     * - `from` and `to` are never both zero.
     *
     * To learn more about hooks, head to xref:ROOT:extending-contracts.adoc#using-hooks[Using Hooks].
     */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {}

    /**
     * @dev Hook that is called after any transfer of tokens. This includes
     * minting and burning.
     *
     * Calling conditions:
     *
     * - when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
     * has been transferred to `to`.
     * - when `from` is zero, `amount` tokens have been minted for `to`.
     * - when `to` is zero, `amount` of ``from``'s tokens have been burned.
     * - `from` and `to` are never both zero.
     *
     * To learn more about hooks, head to xref:ROOT:extending-contracts.adoc#using-hooks[Using Hooks].
     */
    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {}

    //TODO add events and make transfers official
    //FIXED
    function addLiquidity(
        address liquidityProvider,
        uint chain1AssetAmount
    ) external returns (bool, uint, uint) {
        require(msg.sender == _overallContract);
        uint256 liquidityPointsToSend = 0;
        if (_totalSupply == 0) {
            liquidityPointsToSend = 10000000000000000000;
            //9036366617225066000
        } else {
            //9993047142951
            //                 3333333333
            if (_isNativeCoin) {
                liquidityPointsToSend =
                    (_totalSupply * chain1AssetAmount) /
                    (address(this).balance - chain1AssetAmount);
            } else {
                //(1000*47.619)/(1000-47.619)
                liquidityPointsToSend =
                    (_totalSupply * chain1AssetAmount) /
                    (_chain1AssetInterface.balanceOf(address(this)) -
                        chain1AssetAmount);
            }
        }
        //_chain1AssetInterface.transferFrom(_overallContract, address(this), c1aAmount);
        require(liquidityPointsToSend != 0, "LP tokens to send is zero");
        _totalSupply += liquidityPointsToSend;
        _balances[liquidityProvider] += liquidityPointsToSend;
        //_balances[address(this)] += lptosend;
        //require(transferFrom(address(this), liqprovider, lptosend), "Unsuccessfully sent");
        return (
            true,
            liquidityPointsToSend,
            _totalSupply - liquidityPointsToSend
        );
    }

    function sendToUser(
        uint64 swapRatio,
        address recipientWallet
    ) external returns (bool) {
        //uint amountSent, uint256 totalAmount
        require(msg.sender == _overallContract);
        //DONE change the formula for pricing
        //TODO find way to lower worst-case int size
        //find out total balance of c1a
        //uint totalc1a;
        //uint amountToSend = _chain1AssetInterface.balanceOf(address(this)) - (totalAmount * _chain1AssetInterface.balanceOf(address(this))) / (amountSent + totalAmount);
        //require(_chain1AssetInterface.transfer(recipientWallet, amountToSend), "Unsuccessfully sent");
        if (_isNativeCoin) {
            //(bool finishresult, ) = recipientWallet.call{value: (address(this).balance * amountSent) / (amountSent + totalAmount)}("");
            (bool finishResult, ) = recipientWallet.call{
                value: (address(this).balance * swapRatio) / (oneQuadrillion)
            }("");
            require(
                finishResult,
                "Error sending native coin to recipient wallet"
            );
        } else {
            //require(_chain1AssetInterface.transfer(recipientWallet, (_chain1AssetInterface.balanceOf(address(this)) * amountSent) / (amountSent + totalAmount)), "Unsuccessfully sent");
            require(
                _chain1AssetInterface.transfer(
                    recipientWallet,
                    (_chain1AssetInterface.balanceOf(address(this)) *
                        swapRatio) / oneQuadrillion
                ),
                "Unsuccessfully sent"
            );
        }
        //require(_chain1AssetInterface.transfer(recipientWallet, (_chain1AssetInterface.balanceOf(address(this)) * amountSent) / totalAmount), "Unsuccessfully sent");
        //make sure math rounds down

        return true;
    }

    function getqem() external view returns (uint) {
        return
            _queueEntryMapping[tx.origin][
                _queueEntryMapping[tx.origin].length - 1
            ].sentLP;
    }

    function removeLiqAddToQueue(
        uint redeemLPAmount,
        address recipientWallet
    ) external returns (bool) {
        require(
            msg.sender == _overallContract || recipientWallet == msg.sender,
            "must only be called from overall contract or from person redeeming"
        );
        require(redeemLPAmount > 0, "Redeem amount must be greater than 0");
        require(
            recipientWallet != address(0),
            "ERC20: transfer from the zero address"
        );
        require(
            _balances[recipientWallet] >= redeemLPAmount,
            "ERC20: redeem amount exceeds b1alance"
        );
        //if there is an existing entry in queue for msg.sender, it is overwritten with the arguments of the most recent removeLiqAddToQueue function invocation
        if (_queueEntryMapping[recipientWallet].length > 0) {
            _queueEntryMapping[recipientWallet][1] = queueEntry(
                block.number,
                redeemLPAmount
            );
        } else {
            _queueEntryMapping[recipientWallet].push(
                queueEntry(block.number, redeemLPAmount)
            );
        }
        return true;
    }

    //DONE: make work with nativecoin
    function removeLiqQueue(address redeemer) external returns (bool) {
        require(
            msg.sender == _overallContract || redeemer == msg.sender,
            "must only be called from overall contract or from person redeeming"
        );
        require(_queueEntryMapping[redeemer].length > 0, "No entry found");
        //done edit mapping var next line
        uint totalChain1Asset;
        if (_isNativeCoin) {
            totalChain1Asset = address(this).balance;
        } else {
            totalChain1Asset = _chain1AssetInterface.balanceOf(address(this));
        }

        uint queueEntryIndex = _queueEntryMapping[redeemer].length - 1;
        uint sentLP = _queueEntryMapping[redeemer][queueEntryIndex].sentLP;
        require(
            _balances[redeemer] >= sentLP,
            "ERC20: redeem amount exceeds b2alance"
        );
        require(
            block.number >
                _queueEntryMapping[redeemer][queueEntryIndex].currentBlockNum +
                    20
        );
        //send n c1a where n = (totalc1a * sentLP)/totalLP
        delete _queueEntryMapping[redeemer];
        uint _amountToSend = 0;
        if (sentLP > 0) {
            uint amountToSend = (totalChain1Asset * sentLP) / _totalSupply;
            _balances[redeemer] -= sentLP;
            _totalSupply -= sentLP;
            if (_isNativeCoin) {
                (bool sent, ) = redeemer.call{value: amountToSend}("");
                require(sent);
            } else {
                require(
                    _chain1AssetInterface.transfer(redeemer, amountToSend),
                    "Failed to send asset"
                );
            }
            _amountToSend = amountToSend;
        }

        //DONE: send amountToSend

        emit Redeem(redeemer, _amountToSend, sentLP);
        return true;
    }

    //DONE: make work with native coin
    function removeLiqBypassQueue(
        address redeemer,
        uint64 swapRatio
    ) external returns (bool) {
        require(
            msg.sender == _overallContract,
            "Can only be called for overall contract"
        );
        uint proportionalLPToSend = (_balances[redeemer] * swapRatio) /
            oneQuadrillion;
        uint totalc1a = _chain1AssetInterface.balanceOf(address(this));
        if (_isNativeCoin) {
            totalc1a = address(this).balance;
        }
        uint _amountToSend = 0;
        if (swapRatio > 0 && proportionalLPToSend > 0) {
            uint amountToSend = (totalc1a * proportionalLPToSend) /
                _totalSupply;
            _balances[redeemer] -= proportionalLPToSend;
            _totalSupply -= proportionalLPToSend;
            if (_isNativeCoin && amountToSend > 0) {
                (bool sent, ) = redeemer.call{value: amountToSend}("");
                require(sent, "Failed to send native coin back");
            } else {
                require(
                    _chain1AssetInterface.transfer(redeemer, amountToSend),
                    "Failed to send asset"
                );
            }

            _amountToSend = amountToSend;
        }
        emit Redeem(redeemer, _amountToSend, proportionalLPToSend);
        return true;
    }

    //function sendTip(uint tip_amount, uint tip_denominator, address swapminer) public returns (bool){
    function sendTip(uint tipRatio, address swapminer) public returns (bool) {
        require(_isTipPool, "Can only send tip if native coin pool");
        //proper equation implemented
        require(msg.sender == _overallContract);
        uint amountToSend;
        bool sent;
        if (_isNativeCoin) {
            amountToSend = (address(this).balance * tipRatio) / oneQuadrillion;
            (sent, ) = swapminer.call{value: amountToSend}("");
        } else {
            sent = _chain1AssetInterface.transfer(
                swapminer,
                (_chain1AssetInterface.balanceOf(address(this)) * tipRatio) /
                    oneQuadrillion
            );
        }

        require(sent, "Failed to send tip");
        return true;
    }

    fallback() external payable {}

    receive() external payable {}
}

//DONE: have address for tokens it holds
//DONE: recieve tokens and send back lp tokens
//DONE: send tokens
//DONE: check if method called by deployer
//DONE: send in proportion of argumenta/argumentb
//DONE: remove liquidity
//DONE: add method that burns LP tokens and queues up asset to be sent back in proportion
//DONE: add method that executes the queue
//NOTNEEDED: add function that returns balance of c1a
//DONE: add bool returns
